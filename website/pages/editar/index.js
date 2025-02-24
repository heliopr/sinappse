const urlParams = new URLSearchParams(window.location.search)

/* <div class="carta">
                    <p>Pergunta:</p>
                    <input class="pergunta-input" autocomplete="off" type="text">
                    <p>Resposta:</p>
                    <input class="resposta-input" autocomplete="off" type="text">
                    <button class="button2 excluir-button">Excluir</button>
                </div>*/

function criarCarta(q, a) {
    const pergunta = $("<input></input>").addClass("pergunta-input").attr("autocomplete", "off").attr("placeholder", "Pergunta").attr("type", "text").val(q)
    const resposta = $("<input></input>").addClass("resposta-input").attr("autocomplete", "off").attr("placeholder", "Resposta").attr("type", "text").val(a)
    const excluir = $("<button></button>").addClass("button2").addClass("excluir-button").text("Excluir")
    const carta = $("<div></div>").addClass("carta").append($("<p></p>").text("Pergunta:")).append(pergunta)
        .append($("<p></p>").text("Resposta:")).append(resposta).append(excluir)

    excluir.on("click", () => {
        carta.remove()
        if ($(".carta").length == 0) {
            $("#cartas").append($("<p></p>").text("Esse deck não possui cartas").attr("id", "naopossui"))
        }
    })
    
    $("#cartas").append(carta)
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

            $("#titulo").text("Editar Deck - " + response.deck.name)

            if (response.deck.cards != null && response.deck.cards.length > 0) {
                for (const card of response.deck.cards) {
                    //lastId++
                    criarCarta(card.q, card.a)
                    //cards[lastId] = card
                }

                //console.log(cards)
            }
            else {
                $("#cartas").append($("<p></p>").text("Esse deck não possui cartas").attr("id", "naopossui"))
            }
        }, error: () =>{
            window.alert("um erro ocorreu ao tentar carregar cartas")
        }
    })

    $("#adicionar-button").on("click", () => {
        criarCarta("", "")
        $("#naopossui").remove()
    })

    $("#salvar-button").on("click", () => {
        const cards = []

        $(".carta").each((i, e) => {
            cards.push({
                q: $(e).children(".pergunta-input").val(),
                a: $(e).children(".resposta-input").val()
            })
        })

        $.ajax({
            type: "post",
            url: "/api/decks/"+id+"/cards",
            data: JSON.stringify({
                cards: cards
            }),
            dataType: "json",
            headers: {
                Authorization: authToken
            },
            success: (response) => {
                if (!response.success) {
                    window.alert("não foi possível atualizar deck: " + response.message)
                    return
                }

                window.alert("deck salvo com sucesso!")
            },
            error: () => {
                window.alert("não foi possível atualizar deck: erro desconhecido")
            }
        })
    })
}