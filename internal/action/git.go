package action

import (
	"os"

	memfs "github.com/go-git/go-billy/v5/memfs"
	git "github.com/go-git/go-git/v5"
	http "github.com/go-git/go-git/v5/plumbing/transport/http"
	memory "github.com/go-git/go-git/v5/storage/memory"
	log "github.com/sirupsen/logrus"
)

func cloneRepositorty(repoURL string, branch string, auth http.AuthMethod) (*git.Worktree, error) {
	log.Debug("- clone repository")
	repo, err := git.Clone(memory.NewStorage(), memfs.New(), &git.CloneOptions{
		URL:      repoURL,
		Auth:     auth,
		Progress: os.Stdout,
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
		log.Debugf("- pull branch %s", branch)
		w.Pull(&git.PullOptions{
			SingleBranch: true,
			RemoteName:   branch,
		})
	}
	return w, nil
}

func fetchFilesContent(w *git.Worktree) []repoFile {
	log.Debug("- iterate over the files in the repository")
	return iterateDirs(w.Filesystem, "/")
}

func saveRepository(repoFiles []repoFile, targetDir string, variables Variables) {
	log.Debug("- process files in the repository")
	processRepositoryFiles(targetDir, repoFiles, "/", variables)
}
