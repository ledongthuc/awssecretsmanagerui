// initial state
const state = () => ({
  selectedARN: undefined,
});

// getters
const getters = {};

// actions
const actions = {
  selectARN({ commit }, arn) {
    commit("selectARN", arn);
  },
};

// mutations
const mutations = {
  selectARN(state, arn) {
    state.selectedARN = arn;
  },
};

export default {
  namespaced: true,
  state,
  getters,
  actions,
  mutations,
};
