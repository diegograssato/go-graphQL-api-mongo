db = db.getSiblingDB(process.env.MONGO_DATABASE);


db.createUser({
    user: process.env.MONGO_USERNAME,
    pwd: process.env.MONGO_PASSWORD,
    roles: [
        {
            role: 'dbAdmin',
            db: process.env.MONGO_DATABASE,
        }
    ],
});

db.createCollection(process.env.MONGO_COLLECTION)