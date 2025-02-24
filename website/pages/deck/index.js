let cards =  null
let curr = 0

function mostrarCarta() {
    if (curr >= cards.length) {
        $("#secao").hide()
        $("main").append($("<p></p>").text("Esse deck não possui mais cartas!"))
        $("main").append($("#voltar").show())
        return
    }

    $("#pergunta").text(cards[curr].q)
    $("#resposta").text(cards[curr].a).hide()
    $("#resposta-titulo").hide()

    $("#mostrar-button").text("Mostrar Resposta").addClass("button2")
}

onFetch = () => {
    if (userInfo == null) {
        window.location.href = "/"
        return
    }

    const id = window.location.pathname.split("/").pop()
    $.ajax({
        type: "get",
        url: `/api/decks/${id}/`,
        headers: {
            Authorization: authToken
        },
        success: (response) => {
            if (!response.success) {
                window.alert("um erro ocorreu ao tentar carregar cartas: " + response.message)
                return
            }

            $("#titulo").text("Treinar - " + response.deck.name)

            if (response.deck.cards != null && response.deck.cards.length > 0) {
                cards = response.deck.cards
                    .map(v => ({v, rand: Math.random()}))
                    .sort((a, b) => a.rand - b.rand)
                    .map(({v}) => v)
                $("#secao").show()
                mostrarCarta()
            } else {
                $("main").append($("<p></p>").text("Esse deck não possui cartas!"))
                $("main").append($("#voltar").show())
            }
        }, error: () =>{
            window.alert("um erro ocorreu ao tentar carregar cartas")
        }
    })

    $("#prox-button").on("click", () => {
        curr++
        mostrarCarta()
    })

    $("#mostrar-button").on("click", () => {
        if ($("#mostrar-button").hasClass("button2")) {
            $("#mostrar-button").removeClass("button2").text("Esconder Resposta")
            $("#resposta").show()
            $("#resposta-titulo").show()
        } else {
            $("#mostrar-button").addClass("button2").text("Mostrar Resposta")
            $("#resposta").hide()
            $("#resposta-titulo").hide()
        }
    })

    $("#voltar").on("click", () => {
        window.location.href = "/decks"
    })
}