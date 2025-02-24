/*
<li class="deck-element">
    <a href="/decks/1">DECKZAO 2</a>
    <button>Treinar</button>
    <button class="button2">Editar</button>
</li>
*/

function criarDeck(deck) {
    const treinarBtn = $("<button></button>").text("Treinar")
    const editarBtn = $("<button></button>").text("Editar").addClass("button2")
    const a = $("<a></a>").attr("href", "/deck/"+deck.id).text(deck.name)
    $("#decks-list").append($("<li></li>").addClass("deck-element").append(a).append(treinarBtn).append(editarBtn))

    treinarBtn.on("click", () => {
        window.location.href = "/deck/"+deck.id
    })

    editarBtn.on("click", () => {
        window.location.href = "/editar/"+deck.id
    })
}

onFetch = async () => {
    if (userInfo == null) {
        window.location.href = "/login"
        return
    }

    await $.ajax({
        type: "get",
        url: `/api/users/${userInfo.id}/decks`,
        success: (response) => {
            if (!response.success) {
                window.alert("não foi possível encontrar seus decks")
                console.log(response)
                return
            }

            if (response.decks != null && response.decks.length > 0) {
                for (let deck of response.decks) {
                    criarDeck(deck)
                }
            } else {
                $("#decks-list").append($("<p></p>").text("Você não tem nenhum deck! Comece criando algum"))
            }
        },
        error: () => {
            window.alert("não foi possível encontrar seus decks")
        }
    })


    $("#criar-button").on("click", () => {
        $.ajax({
            type: "post",
            url: "/api/decks",
            data: JSON.stringify({
                name: $("#criar-input").val()
            }),
            dataType: "json",
            headers: {
                Authorization: authToken
            },
            success: (response) => {
                if (!response.success) {
                    window.alert("não foi possível criar deck: " + response.message)
                    return
                }

                window.location.href = "/decks"
            },
            error: () => {
                window.alert("não foi possível criar deck: erro desconhecido")
            }
        })
    })
}