import Two from 'two.js';

const container = document.getElementById('game-canvas');

const params = {
    autostart: true,
    height: container.clientHeight,
    width: container.clientWidth,
    type: Two.Types.svg,
};
const two = new Two(params).appendTo(container);

const R = 30;
const r = Math.sqrt(R ** 2 - (R / 2) ** 2)
const sides = 6;
const padding = 5;
const min_l = 15, max_l = 65;
const max_value = 100;


// const w_n = 10;
// const h_n = Math.floor((container.clientHeight - r) / r);
const field = two.makeGroup()

var hexagons = []
for (let i = 0; ; i++) {
    hexagons.push([]);
    const y = padding + r + i * r + i * padding / 2;
    if (y + r + padding >= container.clientHeight) {
        break;  
    }
    for (let j = 0; ; j++) {
        if (Math.random() > 0.3) {
            
            const x = padding + R + j * 3 * R + 2 * padding * j + (i % 2) * padding + R * 1.5 * (i % 2);
            if (x + R + padding >= container.clientWidth) {
                break;  
            }
            console.log(y + R)

            const hexagon = two.makePolygon(0, 0, R, sides);
            hexagon.fill = 'rgba(0, 0, 0, 0.2)';
            hexagon.stroke = '#00a8ff';
            hexagon.linewidth = 1;


            var value = Math.random() > 0.5? Math.ceil(100 * Math.random()): 0;
            const ratio = Math.min(value / max_value, 1);
            const light = max_l - (ratio * (max_l - min_l));
            
            if (value !== 0) {
                hexagon.fill = `hsl(200, 70%, ${light}%)`;
            }
            
            const text = two.makeText(`${value}`, 0, 0);
            text.fill = value !== 0? 'white': 'transparent';
            text.size = 12;
            text.weight = 700;  
            text.family = 'sans-serif';
            
            const group = two.makeGroup();
            
            group.add(hexagon, text);
            group.translation.set(x, y);

            hexagons[i].push(group);
            field.add(group)
        } else {
            hexagons[i].push(null);
        }
    }
}

setupInteractions(two);

function setupInteractions(instance) {
    const svg = instance.renderer.domElement;
    var on = false;

    const getShape = (e) => {
        if (e.target.tagName === 'path') {
            return instance.scene.getById(e.target.id);
        }
        return null;
    };

    svg.addEventListener('mouseover', (e) => {
        const shape = getShape(e);
        if (shape) {
            shape.scale = 1.05; 
        }
    });

    svg.addEventListener('mouseout', (e) => {
        const shape = getShape(e);
        if (shape) {
            shape.scale = 1.0;
        }
    });

    svg.addEventListener('click', (e) => {
        const shape = getShape(e);
        if (shape) {
            if (shape.stroke === '#f5f6fa') {
                shape.stroke = '#00a8ff';
            } else {
                shape.stroke = '#f5f6fa';
            }
            if (shape.linewidth == 4) {
                shape.linewidth = 1;
            } else {
                shape.linewidth = 4;
            }
            on = ~on;
        }
    });
}

two.update();
// console.log(hexagons)
