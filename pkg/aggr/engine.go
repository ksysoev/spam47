package aggr

import (
	"strings"

	"github.com/bbalet/stopwords"
	"github.com/jbrukh/bayesian"
	"github.com/ksysoev/spam47/pkg/repo"
)

type SpamEngine struct {
	classifier *bayesian.Classifier
}

const (
	Spam bayesian.Class = "Spam"
	Ham  bayesian.Class = "Ham"
)

func NewSpamEngine(repo repo.EngineRepo) (*SpamEngine, error) {
	classifier := bayesian.NewClassifier(Spam, Ham)

	spam, err := repo.Load(string(Spam))
	if err != nil {
		return nil, err
	}

	classifier.Learn(spam, Spam)

	ham, err := repo.Load(string(Ham))
	if err != nil {
		return nil, err
	}

	classifier.Learn(ham, Ham)

	return &SpamEngine{
		classifier: classifier,
	}, nil
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

	switch class {
	case "spam":
		e.classifier.Learn(words, Spam)
	case "ham":
		e.classifier.Learn(words, Ham)
	}

	return nil
}
