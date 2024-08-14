import typescript from "rollup-plugin-typescript2";
import commonjs from "@rollup/plugin-commonjs";
import resolve from "@rollup/plugin-node-resolve";
import json from '@rollup/plugin-json';

import { glob } from "glob";

const inputFiles = glob.sync("src/**/*.ts");

export default [
  {
    input: inputFiles,
    output: [
      {
        dir: "dist",
        format: "cjs",
        entryFileNames: "[name].js",
        chunkFileNames: "[name].js",
      },
      {
        dir: "dist",
        format: "esm",
        entryFileNames: "[name].mjs",
        chunkFileNames: "[name].mjs",
      },
    ],
    plugins: [typescript(), resolve(), commonjs(),  json(),],
    external: [],
  },
];
