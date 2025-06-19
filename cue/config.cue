package config 

app: {
    name: string
    port: int &< >0 & <65536
    image: {
        repository: string
        tag:        string
    }
}



//Defaultvärden
app: {
    name: "myapp"
    port: 8080
    image: {
        repository: "sofiamoreta/myapp"
        tag :       "latest"
    }
}