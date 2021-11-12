const down = document.querySelector('.down-screen>span'),
    up = document.querySelector('.up-screen>span')

let equalPressed = false,
    startNewNumber = false

const socket = new WebSocket('ws://127.0.0.1:1234/calculator')
window.onload = () => {
    socket.onopen = () => {
        console.log('Server connected.')
    }
    socket.onmessage = e => {
        down.textContent = e.data
    }
    socket.onerror = () => {
        console.error('Socket error.')
    }
    socket.onclose = e => {
        console.log('Connection closed: ' + e.code)
    }
}

/**
 * 将下半屏的内容加入上半屏。
 */
const downToUp = () => {
    up.textContent += down.textContent
    down.textContent = ''
    startNewNumber = true
}

/**
 * 判断运算符是否为四则运算。
 * @param {string} op 要检验的运算符。
 * @returns 是四则运算为true，不是四则运算为false。
 */
const isArith = op => (/^[+|\-|×|÷]$/).test(op)

/**
 * 判断表达式内是否有数字。
 * @param {string} exp 要检验的表达式。
 * @returns 有数字为true，无数字为false。
 */
const hasDigit = exp => (exp.match(/[\d|π|e]/g) !== null)

/**
 * 判断表达式内是否有小数点。
 * @param {string} exp 要检验的表达式。
 * @returns 有小数点为true，无小数点为false。
 */
const hasDot = exp => (exp.match(/[.|π|e]/g) !== null)

/**
 * 判断表达式内括号是否成对。
 * @param {string} exp 要检验的表达式。
 * @returns 成对为true，不成对为false。
 */
const isParenPaired = exp => {
    let left = exp.match(/\(/g)
    let right = exp.match(/\)/g)
    if (left === null && right === null) {
        return true
    } else if (left === null || right === null) {
        return false
    } else if (left.length != right.length) {
        return false
    }
    return true
}

/**
 * 判断表达式是否有效。
 * @param {string} exp 要检验的表达式。
 * @returns 有效为true，无效为false。
 */
const isValid = exp => hasDigit(exp) && isParenPaired(exp)

/**
 * 将上半屏的字符串转换为可计算的表达式。
 * @returns 发给后端计算的表达式。
 */
const toExpression = () => {
    let text = up.textContent.
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
        down.textContent = '0'
        console.warn('Invalid expression.')
    }
    return text
}

/**
 * 获取用户所按按钮的文本。
 * @param {MouseEvent} e 触发按钮的事件。
 * @returns 按钮上的文本。
 */
const getButtonText = e => {
    const { target } = e
    return target.textContent
}

/**
 * 按等号后重置上半屏。
 */
const resetAfterEqual = () => {
    if (!equalPressed) {
        return
    }
    equalPressed = false
    up.textContent = ''
}

/**
 * 处理四则运算。
 * @param {MouseEvent} e 按钮点击事件。
 */
const onclickArith = e => {
    resetAfterEqual()
    const op = getButtonText(e)
    if (up.textContent == '' && down.textContent == '') {
        down.textContent = '0'
    }
    downToUp()
    if (isArith(up.textContent.slice(-1))) {
        up.textContent = up.textContent.slice(0, -1)
    }
    up.textContent += op
}

/**
 * 处理退格。
 * @param {MouseEvent} e 按钮点击事件。
 */
const onclickBackspace = e => {
    if (down.textContent.length == 0) {
        return
    }
    down.textContent = down.textContent.slice(0, -1)
}

/**
 * 处理清除。
 * @param {MouseEvent} e 按钮点击事件。
 */
const onclickClear = e => {
    up.textContent = ''
    down.textContent = ''
}

/**
 * 处理常数。
 * @param {MouseEvent} e 按钮点击事件。
 */
const onclickConstant = e => {
    resetAfterEqual()
    const constant = getButtonText(e)
    if (down.textContent != '') {
        down.textContent += '×'
        downToUp()
    }
    down.textContent = constant
}

/**
 * 处理数字。
 * @param {MouseEvent} e 按钮点击事件。
 */
const onclickDigit = e => {
    resetAfterEqual()
    const digit = getButtonText(e)
    if (startNewNumber) {
        down.textContent = ''
        startNewNumber = false
    }
    down.textContent += digit
}

/**
 * 处理小数点。
 * @param {MouseEvent} e 按钮点击事件。
 */
const onclickDot = e => {
    resetAfterEqual()
    if (hasDot(down.textContent)) {
        return
    }
    if (down.textContent == '') {
        down.textContent = '0'
    }
    down.textContent += '.'
}

/**
 * 处理正负号。
 * @param {MouseEvent} e 按钮点击事件。
 */
const onclickNegative = e => {
    resetAfterEqual()
    if (down.textContent.startsWith('-')) {
        down.textContent = down.textContent.slice(1)
    } else {
        down.textContent = '-' + down.textContent
    }
}

/**
 * 处理后缀运算符。
 * @param {MouseEvent} e 按钮点击事件。
 */
const onclickPostfix = e => {
    resetAfterEqual()
    const op = getButtonText(e)
    downToUp()
    if (op == '1/x') {
        up.textContent += '^(-1)'
    } else if (op == 'n!') {
        up.textContent += '!'
    } else if (op == 'x²') {
        up.textContent += '²'
    } else if (op == ')') {
        up.textContent += ')'
    }
}

/**
 * 处理前缀运算符。
 * @param {MouseEvent} e 按钮点击事件。
 */
const onclickPrefix = e => {
    resetAfterEqual()
    const op = getButtonText(e)
    if (op == '√x') {
        up.textContent += '√'
    } else {
        up.textContent += op
    }
}

/**
 * 处理等号。
 * @param {MouseEvent} e 按钮点击事件。
 */
const onclickEqual = e => {
    resetAfterEqual()
    downToUp()
    equalPressed = true
    const expression = toExpression()
    console.log(expression)
    socket.send(expression)
}

const arithBtn = document.querySelectorAll('.arith')
arithBtn.forEach(btn => btn.addEventListener('click', e => onclickArith(e)))

const backspaceBtn = document.querySelector('#backspace')
backspaceBtn.addEventListener('click', e => onclickBackspace(e))

const clearBtn = document.querySelector('#clear')
clearBtn.addEventListener('click', e => onclickClear(e))

const constantBtn = document.querySelectorAll('.constant')
constantBtn.forEach(btn => btn.addEventListener('click', e => onclickConstant(e)))

const digitBtn = document.querySelectorAll('.digit')
digitBtn.forEach(btn => {
    btn.addEventListener('click', e => onclickDigit(e))
})

const dotBtn = document.querySelector('#dot')
dotBtn.addEventListener('click', e => onclickDot(e))

const equalBtn = document.querySelector('#equal')
equalBtn.addEventListener('click', e => onclickEqual(e))

const negativeBtn = document.querySelector('#negative')
negativeBtn.addEventListener('click', e => onclickNegative(e))

const postfixBtn = document.querySelectorAll('.postfix')
postfixBtn.forEach(btn => btn.addEventListener('click', e => onclickPostfix(e)))

const prefixBtn = document.querySelectorAll('.prefix')
prefixBtn.forEach(btn => btn.addEventListener('click', e => onclickPrefix(e)))
