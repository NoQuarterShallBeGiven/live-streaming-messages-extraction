const source = "odysee"
const targetEM = '.livestream__comments'; //done  // mutation observer
const cashEM = ".credit-amount p"
const authorEM = ".comment__author"
const textEM = ".livestream-comment__text p" // done


const observer = new MutationObserver(handleMutations);
let socket = new WebSocket("ws://localhost:8839/ws");
socket.onopen = function () {
    console.log("Websocket Status: Connected");
    let count = 0;
    const interval = setInterval(() => {
        if (document.querySelector(targetEM)) {
            clearInterval(interval);
            const targetNode = document.querySelector(targetEM);
            console.log("found ", targetEM, " running chat extractor");
            observer.observe(targetNode, { childList: true });
        } else {
            count++;
            console.log(targetEM, " not found: try", count);
            if ( count === 60 ) {
                console.log("giving up");
                clearInterval(interval);
            }
        }
    }, 1000);
};
socket.onmessage = function (e) {
    console.log("Server Message: " + e.data);
};
socket.onclose = function (e) {
    observer.disconnect();
    console.log('socket closed try again');
    console.log(e);
}
function handleMutations(mutationsList, observer) {
    mutationsList.forEach((mutation) => {
        if (mutation.type === 'childList') {
            mutation.addedNodes.forEach((addedNode) => {
                let cash = "";
                try {
                    cash = addedNode.querySelector(cashEM).textContent.trim();
                }
                catch (e) {
                    console.log("Failed to get cash amount", e);
                }
                let user = "";
                try {
                    user = addedNode.querySelector(authorEM).textContent.trim();
                }
                catch (e) {
                    console.log("Failed to get target username", e);
                }
                let text = "";
                try {
                    text = addedNode.querySelector(textEM).textContent.trim();
                }
                catch (e) {
                    console.log("Failed to get target comment", e);
                }
                // TODO : add source
                socket.send(JSON.stringify({Source: source, User: user, Comment: text, Amount: cash}))
            });
        }
    });
}
