import initGrid from 'main.js'
import Two from "two.js"

function calc_radius() {
    var le = 0, ri = Math.max(container.clientHeight, container.clientWidth);
    while (le + 1 < ri) {
        var mid = (le + ri) / 2;
        const radius = Math.sqrt(mid ** 2 - (mid / 2) ** 2);
        const y = padding + radius + h_n * radius + h_n * padding / 2
        const x = padding + mid + (w_n - 1) * 3 * mid + 2 * padding * (w_n - 1) + padding * 2 + mid * 2;
        if (x <= container.clientWidth && y <= container.clientHeight) {
            le = mid;
        } else {
            ri = mid;
        }
    }
    return le;
}

const container = document.getElementById('game-canvas');
const params = {
    autostart: true,
    height: container.clientHeight,
    width: container.clientWidth,
    type: Two.Types.svg,
};
const two = new Two(params).appendTo(container);

const w_n = 10;
const h_n = 20;
const padding = 8;  
const R = calc_radius();
const r = Math.sqrt(R ** 2 - (R / 2) ** 2);
const sides = 6;
const min_l = 20, max_l = 50;
const max_value = 8;
const text_size = r * 0.8;


initGrid()