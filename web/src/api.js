import axios from 'axios';

const uriPrefix = "/api/v1";

const Api = {

    LogInUser: async function(email, password) {
        const url = uriPrefix + "/login";
        const response = await axios.post(url, {username: email, password:password});
        return JSON.parse(response.data);
    },

    GetUser: async function(userId, jwtToken) {
        const url = uriPrefix + "/users/" + userId;
        const config = {headers: {Authorization: "Bearer " + jwtToken}};
        const response = await axios.get(url, config);
        return response.data;
    },

};

export default Api;



