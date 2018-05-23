import axios from 'axios';

const Api = {
    LogInUser: async function(email, password) {
        const url = "/api/v1/login";
        const response = await axios.post(url, {username: email, password:password});
        const token = response.data;
        console.log("Got token: " + token);
        return token;
    }
};

export default Api;



