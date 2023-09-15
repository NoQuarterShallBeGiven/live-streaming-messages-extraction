const services = ["odysee", "rumble", "gtv"]
let body = document.body;
let menu = document.createElement("div");
menu.style.border = "solid red 1px";
menu.style.backgroundColor = "black";
menu.style.position = "absolute";
menu.style.top = "0";
menu.style.left = "0";
menu.style.zIndex = "2147483646";
body.appendChild(menu);
let connectionStatus = document.createElement("p")
let scriptStatus = document.createElement("p")
let extractor = (s, t, c, a, x) => {
    let socket = new WebSocket("ws://localhost:8839/ws");
    let handleMutations = (mutationsList, observer) => {
        mutationsList.forEach((mutation) => {
            if (mutation.type === 'childList') {
                mutation.addedNodes.forEach((addedNode) => {
                    let cash = "", user = "", text = "";
                    try {
                        cash = addedNode.querySelector(c).textContent.trim();
                    }
                    catch (e) {
                        //console.log("Failed to get cash amount", e);
                    }
                    try {
                        user = addedNode.querySelector(a).textContent.trim();
                    }
                    catch (e) {
                        //console.log("Failed to get target username", e);
                    }
                    try {
                        text = addedNode.querySelector(x).textContent.trim();
                    }
                    catch (e) {
                        //console.log("Failed to get target comment", e);
                    }
                    // TODO : add source
                    socket.send(JSON.stringify({Source: s, User: user, Comment: text, Amount: cash}))
                });
            }
        });
    }
    const observer = new MutationObserver(handleMutations);
    socket.onopen = () => {
        console.log("Websocket Status: Connected");
        connectionStatus.innerText = "Connected";
        connectionStatus.style.color = "green";
        let count = 0;
        const interval = setInterval(() => {
            if (document.querySelector(t)) {
                clearInterval(interval);
                const targetNode = document.querySelector(t);
                console.log("found ", t, " running chat extractor");
                scriptStatus.style.color = "green";
                scriptStatus.innerText = "Running"
                observer.observe(targetNode, { childList: true });
            } else {
                count++;
                console.log(t, " not found: try", count);
                if ( count === 60 ) {
                    console.log("giving up");
                    clearInterval(interval);
                }
            }
        }, 1000);
    };
    socket.onmessage = (e) => {
        console.log("Server Message: " + e.data);
    };
    socket.onclose =  (e) => {
        observer.disconnect();
        connectionStatus.innerText = "Disconnected";
        connectionStatus.style.color = "red";
        scriptStatus.style.color = "red";
        scriptStatus.innerText = "Stopped"
        console.log('socket closed try again');
        console.log(e);
    }
}
let addEm = () => {
    let insertBreak = (em) => {
        em.appendChild(document.createElement("br"))
    };
    let dragElement = (em) => {
        let pos1 = 0, pos2 = 0, pos3 = 0, pos4 = 0;
        em.onmousedown = (e) => {
            pos3 = e.clientX;
            pos4 = e.clientY;
            document.onmouseup = () => {
                document.onmouseup = null;
                document.onmousemove = null;
            };
            document.onmousemove = (e) => {
                e.preventDefault();
                pos1 = pos3 - e.clientX;
                pos2 = pos4 - e.clientY;
                pos3 = e.clientX;
                pos4 = e.clientY;
                em.style.top = (em.offsetTop - pos2) + "px";
                em.style.left = (em.offsetLeft - pos1) + "px";
            };
        }
    }
    let addOpt = (em, opt) => {
        let option = document.createElement("option");
        option.value = opt;
        option.text = opt;
        em.appendChild(option);
    };
    connectionStatus.style.color = "red";
    connectionStatus.innerText = "Not Connected"
    connectionStatus.style.margin = "1px";
    scriptStatus.style.color = "red";
    scriptStatus.innerText = "Not running"
    scriptStatus.style.margin = "1px";
    menu.appendChild(connectionStatus);
    menu.appendChild(scriptStatus);
    let select = document.createElement("select");
    select.style.border = "solid red 1px";
    select.style.backgroundColor = "black";
    select.style.color = "red";
    services.forEach((value)=>{
        addOpt(select, value);
    })
    menu.appendChild(select);
    const button = document.createElement('button');
    button.textContent = 'start extractor';
    button.addEventListener('click', () => {
        start(select.value);
    });
    button.style.zIndex = "2147483647";
    button.style.color = "red";
    button.style.backgroundColor = "black";
    button.style.border = "solid red 1px";
    insertBreak(menu);
    menu.appendChild(button);
    dragElement(menu);
    let start = (service) => {
        console.log("loading", select.value);
        switch (select.value) {
            case "odysee":
                extractor("odysee", ".livestream__comments", ".credit-amount p", ".comment__author", ".livestream-comment__text p")
                break;
            case "rumble":
                extractor("rumble", ".chat-history ul", ".credit-amount p",".chat-history--username a", ".chat-history--message")
                break;
            case "gtv":
                extractor("gtv", "#room-messages", "TODO", ".message-username span", ".message-text span")
                break;
        }
    };
    let close = () => {
        // TODO
        console.log("removing menu");
        menu.remove();
    };
};
addEm();
