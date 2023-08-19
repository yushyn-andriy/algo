#!/bin/bash

# Create the directory structure
mkdir foo_package
cd foo_package
mkdir foo

# Create the foo_module.py file with content
echo 'def hello():
    return "Hello from foo!"' > foo/foo_module.py

# Create __init__.py (empty for simplicity)
touch foo/__init__.py

# Create the setup.py file
echo 'from setuptools import setup

setup()' > setup.py

# Create the setup.cfg file
echo '[metadata]
name = foo-package
version = 0.1
description = A simple example package named foo

[options]
packages = find:' > setup.cfg

# Build the package
python setup.py sdist

# Optionally build a wheel (requires wheel package)
# python setup.py bdist_wheel

# Echo the completion
echo "Packaging completed! Check the 'dist' directory."

