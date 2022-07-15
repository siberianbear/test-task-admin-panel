<script lang="ts">

  import {apiGetRequest, apiDeleteRequest, apiPutRequest} from "./Utils/API"
  import SideBar from './components/SideBar.svelte'
  import UsersList from "./components/UsersList.svelte"
  import Modal from "./components/Modal.svelte"

  let logo = '/imgs/logo-solana.png'

  // variable user for to clean active (selected) user in user list, and that's reactive - that's why it's in classes of children
  let cleanActiveUser = false;

  let stepOneValidation = async (e:string): Promise<Object> => {
    const { response } = await apiGetRequest({
        path: 'getallusers',
      });

    let res: JSON = JSON.parse(response)

    return res;

  }

  // да, надо все интерфейсы в отдельный файлик, но не успел :(
  interface User {
    uid: number;
    name: string;
    email: string;
    phone: string;
    gender: string;
  }

  let ready_data: Promise<Object> = stepOneValidation("")

  let activateControls: boolean = false;
  let currentUser: User = {} as User;

  let isModalOpened: boolean = false;

  let activateButtons = (e: Object): void => {
    activateControls = true;
    currentUser = e.detail.user
  }

  let signalFromModal = (e: Object): void => {

    switch (e.detail.action) {
      case "updateUser":
        currentUser = e.detail.user
        activateControls = false;
        cleanActiveUser = true;
        setTimeout(()=>{
          cleanActiveUser = false;
          isModalOpened = false;
          ready_data = stepOneValidation("")
        }, 200)

        // console.log("going to update user", currentUser)

        const updateUser = async (user_Id: Number): Promise<void> => {
          const {response, apiError } = await	apiPutRequest({
              path: 'edituser/' + user_Id,
              data: currentUser
            });

          console.log(apiError)

          return response;
        }
        updateUser(e.detail.user.uid);
        
        break;
    
      case "cancel":
        activateControls = false;
        cleanActiveUser = true;
        
        currentUser = {uid: undefined, name: '', email: '', phone: '', gender: ''};

        setTimeout(()=>{
          cleanActiveUser = false;
        }, 100)
        isModalOpened = false;
        break;

      default:
        break;
    }

  }

  let sidebarAction = (e: Object): void => {

    switch (e.detail.action) {
      case "edit":
        currentUser = e.detail.user
        isModalOpened = true;

        break;
    
      case "remove":
        
        const deleteUser = async (user_Id: Number): Promise<void> => {
          const { request, response, apiError } = await	apiDeleteRequest({
              path: 'removeuser/' + user_Id,
              data: 'test'
            });

          // console.log(apiError)
          return response;
          // console.log(request);
        }
        deleteUser(e.detail.user.uid);
        activateControls = false;

        setTimeout(()=>{
          ready_data = stepOneValidation("")
        }, 300)

        break;

      default:
        break;
    }
  }

</script>

<main>

  <Modal {isModalOpened} {currentUser} on:modalAction={signalFromModal} />

  <header>
    <div class="logo">
      <img src={logo} alt="">
    </div>
    <div>
      <div class="main-title">Test Admin Panel</div>
    </div>
  </header>

  <div class="wrapper">

    <div class="sidebar">
      <SideBar {activateControls} {currentUser} on:sidebarAction={sidebarAction} />
    </div>
    <div class="content">
      <UsersList on:selected={activateButtons} {ready_data} {cleanActiveUser} />
    </div>
    
  </div>

</main>

<style>
  :root {
    font-family: -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, Oxygen,
      Ubuntu, Cantarell, 'Open Sans', 'Helvetica Neue', sans-serif;
  }

  main {
    margin: 0 !important;
    padding: 0 !important;
    height: 100vh;
  }

  .logo img {
    width: 9rem;
  }

  header {
    background-color: rgba(0, 0, 0, 0.745);
    padding: 1rem 2rem;
    display: flex;
    align-items: center;
    justify-content: space-between;
    margin-bottom: 10px;
  }

  .wrapper {
    display: flex;
    flex-direction: column;
    background-color: rgba(255,255,255, .7);
    position: relative;
  }
  

  main {
    background-image: url("./imgs/farawaybg.png");
    background-repeat: no-repeat;
    background-position: center;
    background-size: contain;
    background-position: bottom;
  }

  .sidebar {
    display: flex;
    justify-content: center;
    padding: 1rem;
    background-color: rgba(172, 172, 172, 0.573);
  }

  .content {
    padding: 1rem 0;
  }

  .main-title {
    display: none;
    color: #ff3e00;
    color: white;
    text-transform: uppercase;
    font-size: 1.6rem;
    font-weight: 100;
    text-align: center;
    margin: 0;
  }

  @media (min-width: 480px) {
    .sidebar {
      justify-content: flex-start;
    }

    .content {
      padding: 1rem;
    }

  }

  @media (min-width: 520px) {
    .main-title {
      display: block;
    }

  }

  @media (min-width: 860px) {
    
    .wrapper {
      flex-direction: row;
    }
  }

</style>
