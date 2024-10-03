import (
    "errors"
    "github.com/patrcoff/la-tool/src/cli/commands"
)

httpsUrls := []string{
    "https://github.com/username/repository.git",
    "https://github.com/organization/project.git",
    "https://gitlab.com/username/repository.git",
    "https://username@bitbucket.org/username/repository.git",
    "https://dev.azure.com/organization/project/_git/repository",
    "https://username@dev.azure.com/organization/project/_git/repository",
    "https://git.example.com/scm/project/repository.git"
}


func TestHttp(t *testing.T) {
    result := make([]bool, 7)
    for (i := 0; i < 7; i++) {
        result.append(commands.DetectSourceType(httpUrls[i]))
    }
    if result != []bool{true, true, true, true, true, true, true} {
        t.Errorf("HTTPS Source type not recognised.")
    }
    // I'd like to extend this to output the ones which fail
}
