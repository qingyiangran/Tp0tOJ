import Vue from "vue";
import Vuex from "vuex";
import GlobalStateStore from "@/stores/GlobalState";
import CompetitionStore from "@/stores/CompetitionState";
import BulletinStore from "@/stores/BulletinState";

Vue.use(Vuex);

export default new Vuex.Store({
  state: {},
  mutations: {},
  actions: {},
  modules: {
    global: GlobalStateStore,
    competition: CompetitionStore,
    bulletin: BulletinStore
  }
});
