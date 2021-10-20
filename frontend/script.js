/* eslint-disable max-lines-per-function */
/* eslint-disable max-statements */
/* eslint-disable complexity */
const buttons = document.querySelectorAll('button'),
    down = document.querySelector('.down-screen>span'),
    up = document.querySelector('.up-screen>span')

let sock = null
const server = 'ws://127.0.0.1:1234/calculator'
window.onload = () => {
    sock = new WebSocket(server)
    sock.onopen = () => {
        console.log('Connection success: ' + server)
    }
    sock.onclose = e => {
        console.log('Connection closed: ' + e.code)
    }
    sock.onmessage = e => {
        console.log('Received message: ' + e.data)
        down.textContent = e.data
    }
}

let equalPressed = false

/**
 * @description Replace content of up-screen with down-screen.
 */
const downToUp = () => {
    if (up.textContent == '0') {
        up.textContent = ''
    }
    if ((/^0[+\-×÷]/).test(down.textContent)) {
        down.textContent = down.textContent.slice(1)
    }
    up.textContent += down.textContent
    down.textContent = '0'
}

/**
 * @description Examine the validity of up-screen expression.
 * @param {String} str The expression to be examined.
 * @returns true if valid, false if invalid.
 */
const isValid = str => {
    if (str.match(/\d/g) === null) {
        return false
    }
    let left = str.match(/\(/g)
    let right = str.match(/\)/g)
    if (left === null) {
        if (right === null) {
            return true
        }
        return false
    }
    if (right === null) {
        return false
    } else if (left.length == right.length) {
        return true
    }
    return false
}

/*
 * Add event listener for every button.
 * This piece of code is shit but it works anyway.
 */
buttons.forEach(button => {
    button.addEventListener('click', e => {
        let { target } = e
        let curText = target.textContent

        if (down.textContent == 'Error') {
            down.textContent = '0'
        }

        if (equalPressed) {
            equalPressed = false
            up.textContent = ''
            if ('+-×÷'.includes(curText)) {
                downToUp()
            } else {
                down.textContent = '0'
                up.textContent = '0'
            }
        }

        if (curText == '←') {
            down.textContent = down.textContent.length == 1
                ? '0'
                : down.textContent.slice(0, -1)
        } else if (curText == 'C') {
            up.textContent = '0'
            down.textContent = '0'
        } else if (curText == '.') {
            if (down.textContent.match(/\./g) === null) {
                down.textContent += curText
            }
        } else if ('+-×÷'.includes(curText)) {
            if (down.textContent == '0' && '+-×÷'.includes(up.textContent.slice(-1))) {
                up.textContent = up.textContent.slice(0, -1) + curText
            } else {
                down.textContent += curText
                downToUp()
            }
        } else if (curText == '+/-') {
            if (down.textContent.startsWith('-')) {
                down.textContent = down.textContent.slice(1)
            } else {
                down.textContent = '-' + down.textContent
            }
        } else if (curText == '1/x') {
            down.textContent += '^(-1)'
        } else if (curText == 'x²') {
            down.textContent += '²'
        } else if (curText == '√x') {
            if (down.textContent == '0') {
                down.textContent = ''
            }
            down.textContent += '√'
        } else if (curText == 'n!') {
            down.textContent += '!'
        } else if (curText == '=') {
            let dotSet = down.textContent.
                replace(/π/g, Math.PI).
                replace(/e/g, Math.E).
                match(/\./g)
            if (dotSet !== null && dotSet.length > 1) {
                down.textContent = 'Error'
                return
            }
            downToUp()
            equalPressed = true
            let text = up.textContent
            text = text.
                replace(/×/g, '*').
                replace(/÷/g, '/').
                replace(/sin/g, 's').
                replace(/cos/g, 'c').
                replace(/tan/g, 't').
                replace(/²/g, 'p').
                replace(/√/g, 'r').
                replace(/\^\(-1\)/g, 'u').
                replace(/\(-/g, '(0-').
                replace(/π/g, Math.PI).
                replace(/e/g, Math.E)
            if (text.startsWith('-')) {
                text = '0' + text
            }
            if (!isValid(text)) {
                down.textContent = 'Error'
                return
            }
            console.log(text)
            sock.send(text)
        } else {
            if (down.textContent.replace('-', '') == '0') {
                down.textContent = down.textContent.replace('0', '')
            }
            down.textContent += curText
        }
    })
})
