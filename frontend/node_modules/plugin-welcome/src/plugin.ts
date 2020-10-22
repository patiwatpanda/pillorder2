import { createPlugin } from '@backstage/core';
import WelcomePage from './components/WelcomePage';
import SignIn from './components/SignIn'
import ShowPillorder from './components/WatchVideo/ShowPillorder';


export const plugin = createPlugin({
  id: 'welcome',
  register({ router }) {
    router.registerRoute('/home', WelcomePage);
    router.registerRoute('/show_pillorder', ShowPillorder);
    router.registerRoute('/', SignIn);
  },
});
