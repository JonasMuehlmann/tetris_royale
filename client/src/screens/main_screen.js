import { AnimatePresence, motion } from "framer-motion";
import { withDialogContext } from "../contexts/dialog-context";
import { Screen, useScreens, withScreenContext } from "../contexts/screen-context"
import ErrorScreen from "./404";
import LogInScreen from "./login";
import LobbyScreen from "./lobby";
import FriendsBox from "../components/friends_box";
import TetrisScreen from "./tetris";
import { KeybindsContextProvider } from "../contexts/keybinds-context";

const MainScreen = () => {
  const {
    currentScreen,
    navigate,
  } = useScreens()

  const renderCurrentScreen = () => {
    return (
      <AnimatePresence
        exitBeforeEnter>
        {
          currentScreen?.name === 'login' ?
            (
              <motion.div
                className="w-screen h-screen"
                exit={{ opacity: 0, scale: .35 }}
                transition={{ duration: 1.5, type: 'spring' }}
                key={currentScreen.name} >
                <LogInScreen />
              </motion.div>
            ) :
            currentScreen?.name === 'menu' ?
              (
                <motion.div
                  className="w-screen h-screen"
                  initial={{ opacity: 0, y: -window.innerHeight }}
                  animate={{ opacity: 1, y: 0 }}
                  exit={{ opacity: 0, y: -window.innerHeight }}
                  transition={{ duration: 1.5, type: 'spring', delay: .25 }}
                  key={currentScreen.name}>
                  <LobbyScreen />
                  <FriendsBox />
                </motion.div>
              ) :
              currentScreen?.name === 'tetris' &&
              (
                <motion.div
                  className="w-screen h-screen"
                  initial={{ opacity: 0, scale: 0 }}
                  animate={{ opacity: 1, scale: 1 }}
                  exit={{ opacity: 0, scale: 0 }}
                  transition={{ duration: 1.5, type: 'spring' }}
                  key={currentScreen.name}>
                  <TetrisScreen />
                </motion.div>
              )
        }
      </AnimatePresence>
    )
  }

  return (
    <div className="z-20">
      <KeybindsContextProvider>
        {
          currentScreen ?
            (
              renderCurrentScreen()
            ) :
            (
              <ErrorScreen
                onNavigate={() => navigate(Screen.LogIn)}
              />
            )
        }
      </KeybindsContextProvider>
    </div>
  )
}

export default withScreenContext(MainScreen)