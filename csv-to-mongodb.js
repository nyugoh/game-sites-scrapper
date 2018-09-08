const csv = require('csvtojson')
const fs = require('fs')
const mongodbClient = require('mongodb').MongoClient;
const assert = require('assert');

let mongoDBUri = 'mongodb://localhost:27017';
let path = './mmz4281/';
let games = [];
let filePaths = [];

mongodbClient.connect(mongoDBUri, function (error, client) {
	assert.equal(null, error);
	console.log('Connected to mongodb')

	const db = client.db('football-data');
	const collection = db.collection('games');

	walk() // Get all objects
	for (var i = 0; i < filePaths.length; i++) {
		csv().fromFile(filePaths[i]).then(obj => {
			console.log('Inserting....')
			collection.insertMany(obj, function (error, result) {
				assert.equal(null, error)
				console.log('Inserted ')
			});
		}).catch(error => {
			console.log(error)
		});
	}
});


function walk() {
	let dirs = fs.readdirSync(path)
	for (var i=0; i<dirs.length; i++) {
    	let files = fs.readdirSync(path+dirs[i]);
		for (var j = files.length - 1; j >= 0; j--) {
			filePaths.push(path+dirs[i]+'/'+files[j]);
		}
    }
}

