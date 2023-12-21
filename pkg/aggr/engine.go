package aggr

import (
	"strings"

	"github.com/bbalet/stopwords"
	"github.com/jbrukh/bayesian"
	"github.com/ksysoev/spam47/pkg/repo"
)

type SpamEngine struct {
	classifier *bayesian.Classifier
	repo       repo.ClassifieRepo
}

func NewSpamEngine(repo repo.ClassifieRepo) (*SpamEngine, error) {
	classifier, err := repo.Load()
	if err != nil {
		return nil, err
	}

	return &SpamEngine{
		classifier: classifier,
		repo:       repo,
	}, nil
}

func (e *SpamEngine) Classifier() *bayesian.Classifier {
	return e.classifier
}

func (e *SpamEngine) Check(message, lang string) (string, float64) {
	msg := stopwords.CleanString(message, lang, true)
	words := strings.Fields(msg)

	probs, indx, _ := e.classifier.ProbScores(words)

	class := e.classifier.Classes[indx]

	return string(class), probs[indx]
}

func (e *SpamEngine) Train(message, class, lang string) error {
	msg := stopwords.CleanString(message, lang, true)
	words := strings.Fields(msg)

	return e.repo.Update(e.classifier, class, words)
}
