var stockFetcher = require("stock-fetcher");

module.exports = {
    price: function(request, response) {
        stockFetcher.getPrice(request.params.stock, function(error, price) {
            if (price === undefined) {
                return response.send("There was an error fetching information for the " + request.params.stock + " stock.");
            } else {
                response.set('Content-Type', 'text/html');
                return response.send(new Buffer("₦" + price));
            }
        });
    }
};
