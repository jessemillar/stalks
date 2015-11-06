module.exports = {
    attributes: {
        // Primitive attributes
        firstName: {
            type: 'string',
            defaultsTo: ''
        },
        lastName: {
            type: 'string',
            defaultsTo: ''
        },
        wallet: {
            type: 'integer',
            defaultsTo: "10000"
        },

        // Attribute methods
        getFullName: function() {
            return this.firstName + ' ' + this.lastName;
        },
        isMarried: function() {
            return !!this.spouse;
        }
    }
};
