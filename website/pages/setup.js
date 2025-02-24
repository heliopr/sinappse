$(document).ready(() => {
    $("header").load("/components/header.html")
    $("footer").load("/components/footer.html")
})



const authToken = localStorage.getItem("auth")
let userInfo = null
let onFetch = null

setTimeout(() => {
    if (authToken != "null") {
        $.ajax({
            type: "get",
            url: "/api/users/info?token="+authToken,
            success: function (response) {
                if (response.success) {
                    delete response.success
                    userInfo = response
                }
            },
            error: function () {
                localStorage.setItem("auth", null)
            }
        }).done(() => {
            if (onFetch != null) {
                onFetch()
            }
        })
    } else {
        setTimeout(() => {
            if (onFetch != null) {
                onFetch()
            }
        }, 1)
    }
}, 1);