<script lang="ts">

  export let activateControls: boolean;
  export let currentUser: User;

  import { createEventDispatcher} from 'svelte';
  const dispatch = createEventDispatcher();

  // да, точно надо интерфейсы в отдельный файл! Знаю-знаю :)
  interface User {
    uid: number;
    name: string;
    email: string;
    phone: string;
    gender: string;
  }
  
  function sendSignal(command: string, currentUser: User): void {
    dispatch('sidebarAction', {
      action: command,
      user: currentUser
    });

  }

  const editUser = (): void => {

    sendSignal("edit", currentUser)

  }

  const removeUser = (): void => {

    sendSignal("remove", currentUser)

  }

</script>

<div class="control-sidebar">

  <div class="title">Control panel</div>
  
  <div class={!activateControls ? "sidebar-hint active" : "sidebar-hint"}>
    <p>Please choose any user by click on a card at Users list </p>
  </div>

  <div class={activateControls ? "sidebar-control-block active" : "sidebar-control-block"}>
    
    <button class="sidebar-control-bttn edit" on:click={editUser}>
      Edit User
    </button>
  
    <button class="sidebar-control-bttn delete" on:click={removeUser}>
      Remove User
    </button>
  
  </div>

</div>



<style>
  .control-sidebar {
    max-width: 280px;
  }

  .title {
    font-weight: 700;
    margin-bottom: 1rem;
  }

  .sidebar-control-block {
    visibility: hidden;
    display: flex;
    flex-direction: column;
    justify-content: center;
  }

  .sidebar-control-block.active {
    visibility: visible;
  }

  .sidebar-hint {
    display: none;
  }

  .sidebar-hint.active {
    display: block;
  }

  .sidebar-hint {
    margin-bottom: 10px;
  }

  .sidebar-control-bttn {
    padding: 1em 1.5em;
    color: white;
    border-radius: .5em;
    outline: none;
    width: 130px;
    border: none;
    cursor: pointer;
    margin: 3px;
    font-weight: 700;
  }

  .sidebar-control-bttn.edit {
    background-color: rgb(1, 150, 24);
  }

  .sidebar-control-bttn.delete {
    background-color: rgb(199, 36, 36);
  }

  /* button:active {
    background-color: rgba(255, 62, 0, 0.2);
  } */

  @media (min-width: 480px) {

    .sidebar-control-block {
      flex-direction: row;
      margin: 5px;
    }

  }
  
</style>
