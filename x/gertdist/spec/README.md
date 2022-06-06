<!--
order: 0
title: "gertdist Overview"
parent:
  title: "gertdist"
-->

# `gertdist`

<!-- TOC -->
1. **[Concepts](01_concepts.md)**
2. **[State](02_state.md)**
3. **[Messages](03_messages.md)**
4. **[Events](04_events.md)**
5. **[Params](05_params.md)**
6. **[BeginBlock](06_begin_block.md)**

## Abstract

`x/gertdist` is an implementation of a Cosmos SDK Module that allows for governance controlled minting of coins into a module account. Coins are minted during inflationary periods, which each period have a governance specified APR and duration. This module does not provide functionality for spending or distributing the minted coins.