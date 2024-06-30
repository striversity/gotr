import PocketBase from 'pocketbase';

const pb = new PocketBase('http://127.0.0.1:8090');

// list and filter "example" collection records
const result = await pb.collection('items').getList(1, 20, {});

// authenticate as auth collection record
const userData = await pb.collection('users').authWithPassword('test@example.com', '123456');

// or as super-admin
const adminData = await pb.admins.authWithPassword('test@example.com', '123456');

// and much more...