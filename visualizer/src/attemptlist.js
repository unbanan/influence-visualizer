document.addEventListener('DOMContentLoaded', () => {
    const attempts = [123, 456, 27]; // later from database

    const div = document.getElementById("list");
    const ul = document.createElement('ul');

    div.appendChild(ul);

    for (let attempt of attempts) {
        var s = "/attempts/" + attempt;
        var a = document.createElement('a');
        var li = document.createElement('li');
        li.appendChild(a);
        li.style.textDecoration = "none";
        li.style.listStyleType = "none";
        a.textContent = `Посылка №${attempt}`;
        li.style.color = "white";
        li.style.fontSize = "1.5rem";
        a.href = s;
        ul.appendChild(li);
    }
});