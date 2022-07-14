<script lang="ts">

  export let currentUser: User;
  export let isModalOpened: boolean;

  import { createEventDispatcher } from 'svelte';
  const dispatch = createEventDispatcher();

  interface User {
    uid: number;
    name: string;
    email: string;
    phone: string;
    gender: string;
  }

  interface Command {
    action: string;
    user?: Object;
  }

  function sendSignal(command: string, editedUser: Object): void {

    let adminsInstruction: Command = {} as Command;
    adminsInstruction.action = command;

    switch (command) {
      case "updateUser":
      adminsInstruction.user = editedUser
        break;
      
      default:
        break;
    }

    dispatch('modalAction', adminsInstruction);
  }

  const compareTwoObjects = (object1: User, object2: User): boolean => {
    const keys1 = Object.keys(object1);
    const keys2 = Object.keys(object2);
    if (keys1.length !== keys2.length) {
      return false;
    }
    for (let key of keys1) {
      if (object1[key] !== object2[key]) {
        return false;
      }
    }
    return true;
  }

  let validateEmail = (email: string): boolean => {
 
    const EMAIL_REGEXP = /^(([^<>()[\].,;:\s@"]+(\.[^<>()[\].,;:\s@"]+)*)|(".+"))@(([^<>()[\].,;:\s@"]+\.)+[^<>()[\].,;:\s@"]{2,})$/iu;
    // const EMAIL_REGEXP = /.+@.+\..+/i

    let result: boolean;

    function isValidEmail(value): boolean {
      return EMAIL_REGEXP.test(value);
    }

    if (isValidEmail(email)) {
      result = true;
    } else {
      result = false;
    }

    return result;

  }

  let saveChanges = (e:Event):void => {
    e.preventDefault();

    let updatedUser: User = {} as User;
    updatedUser.uid = currentUser.uid;
    updatedUser.name = document.getElementById("name").value
    updatedUser.email = document.getElementById("email").value
    updatedUser.phone = document.getElementById("phone").value
    updatedUser.gender = document.getElementById("gender").value

    if (!validateEmail(updatedUser.email)) {
      document.getElementById("email").style.borderColor = 'orange'
      document.getElementById("emailLabel").innerHTML = 'Email is not valid'
      return;
    }
    else {
      document.getElementById("email").style.borderColor = 'revert'
      document.getElementById("emailLabel").innerHTML = 'Email'
    }


    if (compareTwoObjects(currentUser, updatedUser)) {
      console.log("nothing was changed/edited, action is cancelled")
      sendSignal("cancel", "")
    }
    else {
      sendSignal("updateUser", updatedUser)
    }
    
  }

  let cancelChanges = (e:Event):void => {
    e.preventDefault(); 
    sendSignal("cancel", "")
  }

</script>

<div class={isModalOpened ? "modal-wrapper active" : "modal-wrapper"}>

  <div class="modal">

    <div class="title">Edit user</div>

    <form on:submit={saveChanges}>

      <div class="form-item">
        <label for="name">Name</label>
        <input id="name" type="text" value={currentUser.name}>
      </div>

      <div class="form-item">
        <label for="email" id="emailLabel">Email</label>
        <input id="email" type="email" value={currentUser.email}>
      </div>

      <div class="form-item">
        <label for="phone">Phone</label>
        <input id="phone" type="text" value={currentUser.phone}>
      </div>

      <div class="form-item">
        <label for="gender">Gender</label>

        {#if currentUser.gender == "male"}
          <select id="gender">
            <option value="male" selected>Male</option>
            <option value="female">Female</option>
          </select>
        {:else}
          <select id="gender">
            <option value="male">Male</option>
            <option value="female" selected>Female</option>
          </select>
        {/if}

      </div>

      <div class="buttons-block">
        <button class="modal-bttn save" on:click={saveChanges}>
          Save Changes
        </button>
      
        <button class="modal-bttn cancel" on:click={cancelChanges}>
          Cancel
        </button>
      </div>

    </form>
  </div>
  
</div>



<style>
  .modal-wrapper {
    position: absolute;

    width: 100%;
    height: 100%;
    background-color: rgba(0,0,0,.7);

    display: none;
    justify-content: center;
    align-items: center;
    
    z-index: 1;

    margin: -8px;
  }

  .modal-wrapper.active {
    display: flex;
  }

  .modal {
    width: 300px;
    padding: 2rem;
    background-color: rgb(236, 219, 219);
    border-radius: 10px;
  }

  .form-item {
    display: flex;
    flex-direction: column;
    margin-top: 7px;
  }

  .form-item input {
    padding: 5px
  }

  .form-item label {
    color: rgba(0,0,0,.7)
  }

  .buttons-block {
    padding-top: 1rem;
    display: flex;
    flex-direction: column;
    justify-content: space-around;
    align-items: center;
  }

  .title {
    font-size: 1.2rem;
    margin-bottom: 20px;
    text-align: center;
  }

  .modal-bttn {
    padding: 1rem;
    color: white;
    border-radius: .5rem;
    outline: none;
    width: 130px;
    border: none;
    cursor: pointer;
    margin: 3px;
    font-weight: 700;
  }

  .modal-bttn.save {
    background-color: rgb(1, 150, 24);
  }

  .modal-bttn.cancel {
    background-color: rgb(119, 119, 119);
  }

  @media (min-width: 380px) {

    .buttons-block {
      flex-direction: unset;
    }

  }
  
</style>
