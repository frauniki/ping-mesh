from setuptools import setup


setup(
    name='ping-mesh-server',
    version='0.0.1a0',
    author='frauniki',
    url='http://github.com/frauniki/ping-mesh',
    packages=['ping_mesh'],
    package_dir={'': 'server'},
    install_requires=(
        'fastapi',
        'uvicorn[standard]',
    ),
)
