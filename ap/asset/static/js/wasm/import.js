'use strict';

import init, { greet } from './wasm_game_of_life.js';

async function run() {
    await init();

    console.log("greet!");
    greet();
}

run();
