module.exports = {
    sell: function(request, response) {
        Portfolio.findOne({
            firstName: 'Stephanie'
        }).exec(function findOneCB(err, found) {
            return response.send(JSON.stringify(found));
        });
    }
};
