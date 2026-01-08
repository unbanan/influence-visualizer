import express from 'express';
const app = express();
const PORT = process.env.PORT;

app.set('view engine', 'ejs');

const db = [
    { uuid: '123', status: 'Passed' },
    { uuid: '456', status: 'Failed' },
];

app.get('/attempts/:uuid', (req, res) => {
    const uuid = req.params.uuid;
    console.log(uuid);
    const attemptData = db.find(item => item.uuid === uuid);
    res.render('attempt', { 
        attempt: attemptData || null 
    });
}); 


app.get('/attempts/', (req, res) => {
    res.redirect("/");
})

app.listen(PORT, () => {
    console.log(
        `Server is running: http://localhost:${PORT}`
    )
})