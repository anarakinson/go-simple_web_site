package server

import(
    "net/http"
)


type secureFileServer struct {
    http.Dir
}

// Function to disable acces to internal fileserver
func (sfs secureFileServer) Open(name string) (result http.File, err error) {
    f, err := sfs.Dir.Open(name)
    if err != nil {
        return nil, err
    }

    fi, err := f.Stat()
    if err != nil {
        return nil, err
    }
    if fi.IsDir() {
        // Return a response that would have been if directory would not exist:
        return sfs.Dir.Open("not-found")
    }
    return f, nil
}
