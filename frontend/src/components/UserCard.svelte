<script lang="ts">

  import { createEventDispatcher } from 'svelte';
  const dispatch = createEventDispatcher();

  export let userData: User;
  export let cleanActiveUser: boolean;

  // ай отдельный файл для интерфейсов нада, насяльникаа-а!
  interface User {
    uid: number;
    name: string;
    email: string;
    phone: string;
    gender: string;
  }

  const onChoose = (e: Event): void => {

    let allCards = document.querySelectorAll(".user-card");
  
    allCards.forEach( card => {
      card.classList.remove("active")  
    });

    e.target.classList.add("active")
    saveData(userData.uid, userData.name, userData.email, userData.phone, userData.gender)

  }

  function saveData(uid: number, name: string, email: string, phone: string, gender: string) {
    dispatch('selected', {
      user: {uid: uid, name: name, email: email, phone: phone, gender: gender}
    });
  }

</script>
  
  <div class='user-card {userData.gender} {cleanActiveUser}' on:click={onChoose}> 
    <div class="name">{userData.name}</div>
    <div>✉️ {userData.email}</div>
    <div>☎️ {userData.phone}</div>
  </div>

<style>
  .user-card {
    display: block;
    padding: 0.7em 1em;
    color: rgb(255, 255, 255);
    border-radius: .5em;
    min-width: 210px;
    max-width: 210px;
    font-variant-numeric: tabular-nums;
    text-align: left;
    cursor: pointer;
    margin: 5px 5px;
  }
  .user-card > div {
    pointer-events: none;
    margin: 5px 0;
    text-overflow: ellipsis;
    white-space: nowrap;
    overflow: hidden;
  }
  .user-card.male {
    background-color: rgba(83, 121, 255, 0.8);
  }

  .user-card.female {
    background-color: rgba(251, 102, 234, 0.8);
  }
  .user-card .name {
    font-weight: 700;
    font-size: 1.2rem;
    width: 100%;
    border-bottom: 1px solid rgba(236, 236, 236, 0.4);
    margin-bottom: 7px;
    padding-bottom: 3px;
  }
  .user-card.active {
    border: 2px solid #ff0000;
    margin: 3px;
  }
  
  </style>
  