import Two from 'two.js';

// import {Field, Cell, Player, TickView, Visualization} from './utils';

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

function get_hsla_color(hue, percent, light, opacity = 1) {
    return `hsla(${hue}, ${percent}%, ${light}%, ${opacity})`
}

// const cells = [[9, 4, 100, 7, "big"], [8, 5, 100, 12, "big"]]

const cells = [
    [3, 3, 100, 1, "small"],
    [3, 5, 100, 1, "small"],
    [5, 3, 100, 1, "small"],
    [5, 5, 100, 1, "small"],
    [7, 2, 100, 1, "small"],
    [8, 3, 100, 1, "small"],
    [9, 3, 100, 1, "small"],
    [10, 4, 100, 1, "small"],
    [9, 4, 100, 1, "small"],
    [10, 5, 100, 1, "small"],
    [9, 5, 100, 1, "small"],
    [8, 6, 100, 1, "small"],
    [7, 6, 100, 1, "small"],
]

const field = two.makeGroup()

for (let i = 0; i < h_n; i++) {
    const y = padding + r + i * r + i * padding / 2;
    for (let j = 0; j < w_n; j++) {
        const x = padding + R + j * 3 * R + 2 * padding * j + (i % 2) * padding + R * 1.5 * (i % 2);

        const hexagon = two.makePolygon(0, 0, R, sides);
        hexagon.fill = get_hsla_color(0, 0, 67, 0.06);
        hexagon.linewidth = 0;

        const inner_hexagon = two.makePolygon(0, 0, r * 0.9, sides);
        inner_hexagon.fill = get_hsla_color(0, 6, 20, 1.00);
        hexagon.id = `hexagon_id=${i}_${j}`;
        inner_hexagon.rotation = Math.PI / 6;
        inner_hexagon.id = `inner_hexagon_id=${i}_${j}`
        inner_hexagon.linewidth = 1;
        two.update();
        inner_hexagon._renderer.elem.classList.add("untouchable");
        
        // var value = 0;
        // var value = Math.random() > 0.5 ? Math.ceil(max_value * Math.random()) : 0;
        
        // const ratio = Math.min(value / max_value, 1);
        // const light = max_l - (ratio * (max_l - min_l));
        const group = two.makeGroup();
        group.add(hexagon, inner_hexagon);
        for (const cell of cells) {
            const value = cell[3];
            if (cell[0] === i && cell[1] === j) {
                if (cell[4] === "big") {
                    inner_hexagon.radius = r + padding / 2;
                }
                if (value !== 0) {
                    const ratio = Math.min(value / max_value, 1);
                    const light = max_l - (ratio * (max_l - min_l));
                    inner_hexagon.fill = get_hsla_color(cell[2], 30, light);
                    
                    const txt = two.makeText(`${value}`, 0, 0);
                    txt.stroke = 'black';
                    txt.linewidth = .5;
                    txt.fill = value !== 0 ? 'white' : 'transparent';
                    txt.size = text_size;
                    txt.weight = 800;
                    txt.family = 'sans-serif';
                    two.update();
                    txt._renderer.elem.classList.add("untouchable");
                    group.add(txt);
                    break;
                }
            }
        }
        group.translation.set(x, y);
        field.add(group)
        two.update();

        field.translation.set(
            two.width / 2 - field.getBoundingClientRect().width / 2,
            two.height / 2 - field.getBoundingClientRect().height / 2
        );
        two.update();
    }
}

setupInteractions(two);

function setupInteractions(instance) {

    var result = new Set()
    var ctrl_pressed = false;
    const svg = instance.renderer.domElement;

    function isUntouchable(element) {
        if (!element || !element.classList) return false;
        
        if (element.classList.contains('untouchable')) {
            return true;
        }
        
        let parent = element.parentElement;
        while (parent && parent !== svg) {
            if (parent.classList && parent.classList.contains('untouchable')) {
                return true;
            }
            parent = parent.parentElement;
        }
        return false;
    }

    const getShape = (e) => {
        if (e.target.tagName === 'path' || e.target.tagName === 'text') {
            if (isUntouchable(e.target)) {
                return null;
            }
            return instance.scene.getById(e.target.id);
        }
        return null;
    };



    function mouse_pressed(shape) { 
        const hasStroke = shape.linewidth > 0 && shape.stroke;
        if (hasStroke) {
            shape.noStroke();
            shape.linewidth = 0;
        } else {
            shape.stroke = '#f5f6fa';
            shape.linewidth = 4;
        }
        two.update();
    }

    svg.addEventListener('click', (e) => {
        const shape = getShape(e);
        if (shape) {
            if (result.has(shape.id)) {
                result.delete(shape.id);
            } else {
                result.add(shape.id);
            }
            console.log(result);
            mouse_pressed(shape);
        }
    });

    function change_fill(shape) {
        var shapes = ['hsla(0, 0%, 67%, 0.50)', 'hsla(0, 0%, 67%, 0.06)'];
        shape.fill = shapes[1 - shapes.indexOf(shape.fill)]
    }

    svg.addEventListener('mouseover', (e) => {
        const shape = getShape(e);
        if (shape) {
            change_fill(shape);
        }
    });
    
    svg.addEventListener('mouseout', (e) => {
        const shape = getShape(e);
        if (shape) {
            change_fill(shape);
        }
    });
}

two.update();
