use('coolcar');

// Create a new index in the collection.

db.getCollection('account').createIndex({
    open_id:1,
},{
    unique:true,
})

db.getCollection('trip')
  .createIndex(
    {
      "trip.accountid":1,
      "trip.status":1,
    }, {
      unique:true,
      partialFilterExpression:{
        "trip.status":1,
      }
    }
  );