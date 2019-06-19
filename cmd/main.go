package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os/exec"
	"strconv"

	"github.com/spotahome/harbor-replication/pkg/log"
	"github.com/spotahome/harbor-replication/schema"
)

const (
	harborFrom         = "XXXXXXXXX"
	harborUserFrom     = "admin"
	harborPasswordFrom = "PPPPPPPPP"
	harborURLFrom      = "http://" + harborUserFrom + ":" + harborPasswordFrom + "@" + harborFrom + "/api/"
	harborTo           = "XXXXXXXXX"
	harborUserTo       = "admin"
	harborPasswordTo   = "PPPPPPPPP"
	harborURLTo        = "http://" + harborUserTo + ":" + harborPasswordTo + "@" + harborTo + "/api/"
)

func getProjects() ([]schema.Project, error) {

	fromConnect := http.Client{}

	req, err := http.NewRequest(http.MethodGet, harborURLFrom+"/projects/", nil)

	resp, err := fromConnect.Do(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("error status received %d", resp.StatusCode)
	}

	defer resp.Body.Close()

	projects := []schema.Project{}

	err = json.NewDecoder(resp.Body).Decode(&projects)

	if err != nil {
		return nil, err
	}

	return projects, err
}

func postProject(project schema.Project, logger log.Logger) error {

	postreq := schema.ProjectReq{
		ProjectName: project.Name,
		Metadata: &schema.ProjectMetadata{
			AutoScan: "true",
			Public:   "false",
		},
	}

	projectJSON, _ := json.Marshal(postreq)
	req, err := http.NewRequest(http.MethodPost, harborURLTo+"/projects/", bytes.NewBuffer(projectJSON))
	req.Header.Set("Content-Type", "application/json")

	toConnect := http.Client{}
	resp, err := toConnect.Do(req)
	if err != nil {
		return err
	}

	if resp.StatusCode != 201 {
		return fmt.Errorf("error creating project %s status received %d", projectJSON, resp.StatusCode)
	}

	return nil
}

func getRepositories(ProjectID int32, logger log.Logger) ([]schema.Repository, error) {
	fromConnect := http.Client{}
	req, err := http.NewRequest(http.MethodGet, harborURLFrom+"/repositories?project_id="+strconv.FormatInt(int64(ProjectID), 10), nil)
	resp, err := fromConnect.Do(req)

	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("error reading repositories  %d", resp.StatusCode)
	}

	defer resp.Body.Close()

	repositories := []schema.Repository{}

	err = json.NewDecoder(resp.Body).Decode(&repositories)

	if err != nil {
		return nil, err
	}

	return repositories, nil

}

func getTags(repositoryName string, logger log.Logger) ([]schema.Tag, error) {

	fromConnect := http.Client{}
	req, err := http.NewRequest(http.MethodGet, harborURLFrom+"/repositories/"+repositoryName+"/tags", nil)
	resp, err := fromConnect.Do(req)

	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("error reading tags  %d", resp.StatusCode)
	}

	defer resp.Body.Close()

	tags := []schema.Tag{}

	err = json.NewDecoder(resp.Body).Decode(&tags)

	if err != nil {
		return nil, err
	}

	return tags, nil

}

func getImages(projectName string, ProjectID int32, logger log.Logger) ([]string, error) {

	var images []string

	repositories, err := getRepositories(ProjectID, logger)

	if err != nil {
		return images, fmt.Errorf("cant get repositories %s", err)
	}

	for _, repository := range repositories {
		tags, err := getTags(repository.Name, logger)
		if err != nil {
			return images, fmt.Errorf("cant get tags %s", err)
		}

		for _, tag := range tags {
			images = append(images, repository.Name+":"+tag.Name)
		}
	}

	return images, nil
}
func copyImage(image string, logger log.Logger) error {

	cmd := exec.Command("/usr/bin/docker", "pull", harborFrom+"/"+image)
	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr

	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("Cant pull image %s %s %s", err, stderr.String(), out.String())
	}

	cmd = exec.Command("/usr/bin/docker", "tag", harborFrom+"/"+image, harborTo+"/"+image)
	cmd.Stdout = &out
	cmd.Stderr = &stderr

	err = cmd.Run()
	if err != nil {
		return fmt.Errorf("Cant tag image %s %s %s", err, stderr.String(), out.String())
	}

	cmd = exec.Command("/usr/bin/docker", "push", harborTo+"/"+image)
	cmd.Stdout = &out
	cmd.Stderr = &stderr

	err = cmd.Run()
	if err != nil {
		return fmt.Errorf("Cant push image %s %s %s", err, stderr.String(), out.String())
	}

	cmd = exec.Command("/usr/bin/docker", "rmi", harborFrom+"/"+image)
	cmd.Stdout = &out
	cmd.Stderr = &stderr

	err = cmd.Run()
	if err != nil {
		logger.Warn("Cant delete local image %s %s %s", err, stderr.String(), out.String())
	}

	cmd = exec.Command("/usr/bin/docker", "rmi", harborTo+"/"+image)
	cmd.Stdout = &out
	cmd.Stderr = &stderr

	err = cmd.Run()
	if err != nil {
		logger.Warn("Cant delete local image %s %s %s", err, stderr.String(), out.String())
	}

	return nil
}

func loginTo() error {
	cmd := exec.Command("/usr/bin/docker", "login", "-u", harborUserTo, "-p", harborPasswordTo, harborTo)
	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr

	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("Cant login in destiny %s %s %s", err, stderr.String(), out.String())
	}
	return nil

}

func main() {
	logger := log.Base()
	logger.Set("debug")

	projects, err := getProjects()

	if err != nil {
		logger.Error(err)
		return
	}

	for _, project := range projects {

		err = postProject(project, logger)

		if err != nil {
			logger.Warn(err)
		} else {
			logger.Debugf("Added project:", project.Name)
		}

		if project.RepoCount > 0 {
			//Have images to copy
			err = loginTo()

			if err != nil {
				logger.Errorf("Cant log: %s", err)
				return
			}

			images, err := getImages(project.Name, project.ProjectId, logger)
			if err != nil {
				logger.Errorf("Cant get images: %s", err)
				return
			}

			for _, image := range images {
				err = copyImage(image, logger)
				if err != nil {
					logger.Errorf("Cant copy image %s: %s ", image, err)
					return
				}

			}
		}
	}

}
