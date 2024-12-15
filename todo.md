
# 0-Shell Project

## Objective
The objective of this project is to create a simple Unix shell from scratch. Through this project, you will explore the core of the Unix system and its API for process creation and synchronization. The shell will allow execution of commands, and will focus on implementing basic Unix-like functionality without advanced features like pipes or redirection.

---

## Features
The shell must support the following commands, implemented from scratch:
- `echo`
- `cd`
- `ls` (including flags: `-l`, `-a`, and `-F`)
- `pwd`
- `cat`
- `cp`
- `rm` (with the flag `-r`)
- `mv`
- `mkdir`
- `exit`

### Additional Requirements
- The shell must handle program interruption (`Ctrl+D`).
- The project must be written in a compiled language (e.g., C, Go, Rust).
- **External binaries cannot be called**; all functionality must be implemented internally.
- Error management is required for invalid commands, missing arguments, and permissions.

---

## Developer Tasks

### **Developer 1(@sefaye): Infrastructure and Base Commands**
- Implement the shell's main loop (prompt `$`, read commands, execute).
- Commands:
  - `exit`: Exit the shell and return control to the parent shell.
  - `echo`: Display provided arguments, matching Unix behavior.
  - `pwd`: Display the current working directory.

### **Developer 2(@yba): Directory Management**
- Implement navigation and directory management:
  - `cd`: Change the working directory, with error handling.
    - Default behavior (`cd` without arguments): Navigate to the user's home directory.
  - `mkdir`: Create a directory, with error management for permissions and existing directories.
  - `ls`: List files in a directory with support for flags `-l`, `-a`, and `-F`.

### **Developer 3(@Adiane): File Management**
- Commands to manage file content and movement:
  - `cat`: Display the contents of a file, with error handling.
  - `cp`: Copy files or directories with proper error management.
  - `mv`: Move files or directories, ensuring errors are handled gracefully.

### **Developer 4(@Belhadj): Cleanup and Advanced Features**
- Commands for removal and final details:
  - `rm`: Remove files or directories, supporting the `-r` flag for recursive deletion.
- Global error handling for the shell (invalid commands, permissions, arguments).
- Manage program interruption (`Ctrl+D`) to exit cleanly.
- Write documentation and ensure all functionality is well-documented.

---

## Collaboration Workflow

1. **Setup:**
   - Create a shared repository with basic project structure.
   - Define coding conventions and configure a `Makefile` or build script for compilation.

2. **Development:**
   - Developers work on separate branches for their assigned features.
   - Perform code reviews weekly to ensure quality and consistency.

3. **Testing:**
   - Regularly test commands for compliance with Unix behavior.
   - Use automated tests for core functionality.

4. **Finalization:**
   - Merge branches, resolve conflicts, and ensure seamless integration.
   - Document usage, examples, and technical details.

---

## Priorities for Implementation
1. **Core Functionality:** Main loop (`$` prompt), `exit`, `pwd`, `echo`, `cd`, `mkdir`.
2. **File and Directory Management:** Commands like `cat`, `cp`, `mv`, `rm`.
3. **Error Handling:** Ensure user-friendly error messages for all commands.
4. **Advanced Features:** Interruption handling (`Ctrl+D`) and detailed error management.

---

## Audit Checklist

### General
- Written in a compiled language.
- Commands are implemented without calling external binaries.

### Functional Tests
1. Does the shell display a `$` prompt and wait for input?
2. Does the shell execute commands only after pressing "Enter"?
3. Does the `exit` command terminate the shell properly?
4. Are `echo`, `pwd`, `cd`, `ls` (with flags), `mkdir`, `cat`, `cp`, `mv`, and `rm` implemented as specified?
5. Does the shell handle errors gracefully (e.g., invalid commands, missing arguments, permission issues)?
6. Does `Ctrl+D` exit the shell cleanly?

---

## Documentation
Ensure the final documentation includes:
- How to compile and run the shell.
- Supported commands and their usage.
- Known limitations and possible extensions.
