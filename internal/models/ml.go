package models

import (
	"encoding/json"
	"os/exec"
)

type Prediction struct {
	Model1 int `json:"model1"`
	Model2 int `json:"model2"`
}

func GetPredictions(input []float64) (Prediction, error) {
	inputJSON, err := json.Marshal(input)
	if err != nil {
		return Prediction{}, err
	}

	cmd := exec.Command("python3", "../../ml/predict.py", string(inputJSON))
	output, err := cmd.Output()
	if err != nil {
		return Prediction{}, err
	}

	var result Prediction
	if err := json.Unmarshal(output, &result); err != nil {
		return Prediction{}, err
	}

	return result, nil
}
