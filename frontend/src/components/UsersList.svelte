<script lang="ts">

import UserCard from "./UserCard.svelte"
import { createEventDispatcher } from 'svelte';

export let ready_data;
export let cleanActiveUser: boolean;

const dispatch = createEventDispatcher();
  
  function sendSignal(e: Event): void {
    dispatch('selected', {
      user: e.detail.user
    });
  }

</script>

<div class="title">
  Users list
</div>

<div class="users-listing {cleanActiveUser}"> 

  {#await ready_data}
    <div class="loader">
      <img src="./imgs/loading.gif" alt="" />
    </div>
  {:then ready_data}

    {#if ready_data}
      {#each ready_data as user (user)}
        <div>
          <UserCard userData={user} on:selected={sendSignal} {cleanActiveUser} />
        </div>
      {/each}

    {:else}
      <div class="oops">
        <p>
          Балуемся, да? Зачем вы удалили всех пользователей 😱 и где нам теперь новых искать?
        </p>
        <p>
          Надо программисту звонить, чтобы он базу из бэкапа перезалил или разработал модуль для добавления новых пользователей в систему 🥺
        </p> 
      </div>
    {/if}

  {:catch err}
    <h5>Error while loading the data: {err}</h5>
    <p>Last time here was an error about: "http: panic serving 127.0.0.1:58635: pq: sorry, too many clients already"</p>
  {/await}

</div>


<style>
  .loader {
    display: flex;
    height: 35vh;
  }

  .loader img {
    width:25px;
    height:25px;
    margin: 0 auto;
  }

  .title {
    font-weight: 700;
    margin-bottom: 1rem;
  }

  .users-listing {
    display: flex;
    flex-wrap: wrap;
  }

  .oops {
    background-color: lightgreen;
    border-radius: 5px;
    padding: 10px 20px;
  }

</style>
