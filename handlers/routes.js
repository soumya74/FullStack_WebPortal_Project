module.exports = function(app){

    app.get('/', (req, res) => res.send('Hello World!'))

    app.get('/login', function(req, res){
        res.send('Hello World 2!')
    });

    //other routes..
}
