import { BrowserRouter, Route, Routes } from 'react-router-dom'
import './App.css'
import LoginPage from './pages/Login'
import { UserProvider } from './contexts/UserContext'
import ConfigPage from './pages/LibraryPage'

function App() {

  return (
    <>
    <UserProvider>
      <BrowserRouter>
        <Routes>
          <Route path="/" element={<LoginPage />}/>
          <Route path="/ui/library" element={<ConfigPage />}/>
        </Routes>
      </BrowserRouter>
    </UserProvider>
    </>
  )
}

export default App
