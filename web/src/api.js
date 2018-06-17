import axios from 'axios';

const uriPrefix = "/api/v1";

function createConfig(jwtToken) {
    return {headers: {Authorization: "Bearer " + jwtToken}};
}

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

    GetAllYears: async function(jwtToken) {
        const url = uriPrefix + "/movies";
        const config = createConfig(jwtToken);
        return await axios.get(url, config)
    },


    GetMovies: async function(year, jwtToken) {
        const url = uriPrefix + "/movies/" + year;
        const config = createConfig(jwtToken);
        console.log(url);
        const rsp = await axios.get(url, config);
        console.log(rsp);
        return rsp.data
    },

    GetAllMovies: async function(jwtToken) {
        const response = this.GetAllYears(jwtToken);
        const config = createConfig(jwtToken);
        let allMovies = await Promise.all(response.data.map(year => {
            const y = year.Year;
            return axios.get(uriPrefix + year.Uri, config)
                .then(movieRsp => movieRsp.data)
                .then(array => {
                    return {year:y, array: array}
                })
        }));
        let movies = {};
        allMovies.forEach(yearList =>
            movies[yearList.year] = yearList.array);
        return movies;
    },

};

export default Api;



