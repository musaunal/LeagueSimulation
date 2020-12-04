import apisauce from 'apisauce';

const api = apisauce.create({baseURL: 'http://localhost:80/'});       //http://192.168.0.106

export default {
    league: {
        list: async () => {
            return api.get('leagues')
        },
        get: async (id) => {
            return api.get(`leagues/${id}`)
        },
        create: async ({name}) => {
            return api.post('leagues', {name})
        },
        update: async (id, {name}) => {
            return api.put(`leagues/${id}`, {name})
        },
        delete: async (id) => {
            return api.delete(`leagues/${id}`)
        },
    },
    team: {
        list: async (leagueId) => {
            return api.get(`leagues/${leagueId}/teams`)
        },
        create: async (leagueId, {name}) => {
            return api.post(`leagues/${leagueId}/teams`, {name})
        },
        update: async (leagueId, id, {name}) => {
            return api.put(`leagues/${leagueId}/teams/${id}`, {name})
        },
        delete: async (leagueId, id) => {
            return api.delete(`leagues/${leagueId}/teams/${id}`)
        },
    },
    match: {
        distributeFixture: async (leagueId) => {
            return api.post(`leagues/${leagueId}/distribute-fixture`, {})
        },
        playOneWeek: async (leagueId) => {
            return api.post(`leagues/${leagueId}/play-one-week`, {})
        },
        playAll: async (leagueId) => {
            return api.post(`leagues/${leagueId}/play-all`, {})
        },
        predict: async (leagueId) => {
            return api.post(`leagues/${leagueId}/predict-leaders`, {})
        },

    }
}
