// initial state
// shape: [{ id, quantity }]
const state = () => ({
  items: [],
  visibility: "all",
});

// getters
const getters = {
  visibility: (state, getters) => {
    return state.visibility;
  },
};

// // actions
// const actions = {
//   checkout({ commit, state }, products) {
//     const savedCartItems = [...state.items];
//     commit("setCheckoutStatus", null);
//     // empty cart
//     commit("setCartItems", { items: [] });
//     shop.buyProducts(
//       products,
//       () => commit("setCheckoutStatus", "successful"),
//       () => {
//         commit("setCheckoutStatus", "failed");
//         // rollback to the cart saved before sending the request
//         commit("setCartItems", { items: savedCartItems });
//       }
//     );
//   },
// };

// mutations
const mutations = {
  setVisibility(state, visibility) {
    state.visibility = visibility;
  },
};

export default {
  namespaced: true,
  state,
  getters,
  //   actions,
  mutations,
};
