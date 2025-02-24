const searchParams = new URLSearchParams(window.location.search)

const token = searchParams.get("token")
if (token != null) {
    localStorage.setItem("auth", token)
    setTimeout(() => {
        window.location.href = "/"
    }, 1);
}

/*onFetch = () => {
    console.log(userInfo)
}*/