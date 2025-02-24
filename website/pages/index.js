console.log("executou")

onFetch = () => {
    if (userInfo !== null) {
        $("#bemvindo-label").text(`Bem-vindo, ${userInfo.username}!`)
        $("#bemvindo").show()
        console.log("ho")
    }
    
    $("#logout-button").on("click", () => {
        if (userInfo != null) {
            localStorage.setItem("auth", null)
            window.location.href = "/"
        }
    })
}