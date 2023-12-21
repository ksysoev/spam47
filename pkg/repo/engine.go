package repo

import (
	"errors"
	"os"

	"github.com/jbrukh/bayesian"
)

type ClassifieRepo interface {
	Save(classifier *bayesian.Classifier) error
	Load() (*bayesian.Classifier, error)
	Update(classifier *bayesian.Classifier, class string, words []string) error
}

const (
	Spam bayesian.Class = "Spam"
	Ham  bayesian.Class = "Ham"
)

type ClassifierFileRepo struct {
	filepath string
}

func NewClassifierFileRepo(filepath string) *ClassifierFileRepo {
	return &ClassifierFileRepo{
		filepath: filepath,
	}
}

func (r *ClassifierFileRepo) Save(classifier *bayesian.Classifier) error {
	return classifier.WriteToFile(r.filepath)
}

func (r *ClassifierFileRepo) Load() (*bayesian.Classifier, error) {
	if _, err := os.Stat(r.filepath); os.IsNotExist(err) {
		classifier := bayesian.NewClassifier(Spam, Ham)
		if err := r.Save(classifier); err != nil {
			return nil, err
		}

		return classifier, nil
	}

	return bayesian.NewClassifierFromFile(r.filepath)
}

func (r *ClassifierFileRepo) Update(classifier *bayesian.Classifier, class string, words []string) error {
	switch class {
	case "spam":
		classifier.Learn(words, Spam)
	case "ham":
		classifier.Learn(words, Ham)
	default:
		return errors.New("unknown class " + class)
	}

	return r.Save(classifier)
}
