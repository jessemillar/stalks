module.exports = {
    get: function(request, response) {
        return response.send("You have no stocks.");
    },
    setup: function(request, response) {
        // Stuff
    }
};
