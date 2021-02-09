<template>
  <div>
    <div class="sp-container">
      <sp-sign-in />
      <sp-bank-balances />
      <sp-token-send />
      <!-- this line is used by starport scaffolding # 4 -->
		<sp-type-form type="voucher" :fields="['group', 'owner', ]" module="Themis" />
		<sp-type-form type="group" :fields="['name', 'vouchers', 'funds', ]" module="Themis" />
    </div>
  </div>
</template>

<script>
import * as sp from "@tendermint/vue";

export default {
  components: { ...sp },
};
</script>
