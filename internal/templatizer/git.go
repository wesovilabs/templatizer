package templatizer

import (
	"os"

	memfs "github.com/go-git/go-billy/v5/memfs"
	git "github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	http "github.com/go-git/go-git/v5/plumbing/transport/http"
	memory "github.com/go-git/go-git/v5/storage/memory"
	log "github.com/sirupsen/logrus"
)

func cloneRepositorty(repoURL string, branch string, auth http.AuthMethod) (*git.Worktree, error) {
	log.Debug("- clone repository")
	repo, err := git.Clone(memory.NewStorage(), memfs.New(), &git.CloneOptions{
		URL:             repoURL,
		Auth:            auth,
		Progress:        os.Stdout,
		InsecureSkipTLS: true,
	})
	if err != nil {
		log.Error("unexpected error while clonning the repository")
		return nil, err
	}
	w, err := repo.Worktree()
	if err != nil {
		return nil, err
	}
	if branch != "" {
		log.Infof("- checkout branch %s", branch)
		if err := w.Checkout(&git.CheckoutOptions{
			Branch: plumbing.NewRemoteReferenceName("origin", branch),
		}); err != nil {
			return nil, err
		}
	}
	return w, nil
}
