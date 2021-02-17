<template>
  <div>
    <div class="sp-container">
      <sp-sign-in />
      <sp-bank-balances />
      <sp-token-send />
      <!-- this line is used by starport scaffolding # 4 -->
		<sp-type-form path="uprm-inso-4101-2020-2021-s2.Themis.Themis" type="vote" :fields="[ ['creator', 1, 'string'] , ['account', 2, 'string'] , ['poll', 3, 'string'] , ['option', 4, 'string'] ]" />
		<sp-type-form path="uprm-inso-4101-2020-2021-s2.Themis.Themis" type="poll" :fields="[ ['creator', 1, 'string'] , ['group', 2, 'string'] , ['title', 3, 'string'] , ['description', 4, 'string'] , ['options', 5, 'string'] , ['deadline', 6, 'string'] ]" />
		<sp-type-form path="uprm-inso-4101-2020-2021-s2.Themis.Themis" type="account" :fields="[ ['creator', 1, 'string'] , ['group', 2, 'string'] , ['vouchers', 3, 'string'] ]" />
		<sp-type-form path="uprm-inso-4101-2020-2021-s2.Themis.Themis" type="group" :fields="[ ['creator', 1, 'string'] , ['name', 2, 'string'] ]" />
    </div>
  </div>
</template>

<script>
import * as sp from "@tendermint/vue";

export default {
  components: { ...sp },
};
</script>
