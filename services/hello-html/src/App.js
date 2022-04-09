import { useEffect, useState } from 'react'

const App = () => {
  const [uppercase, setUppercase] = useState('hello-html');

  useEffect(() => {
    fetch(`/gecko/uppercase/${uppercase}`)
      .then(res => res.text())
      .then(data => setUppercase(data))
  }, [])

  return (
    <div>
      {uppercase}
    </div>
  );
}

export default App;
