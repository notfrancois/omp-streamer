#include "omp.h"

#ifdef __cplusplus
extern "C" {
#endif

/**
 * Creates a dynamic object in the world
 * 
 * @param modelid Model ID of the object
 * @param x X coordinate
 * @param y Y coordinate
 * @param z Z coordinate
 * @param rx X rotation
 * @param ry Y rotation
 * @param rz Z rotation
 * @param worldid World ID (-1 for all worlds)
 * @param interiorid Interior ID (-1 for all interiors)
 * @param playerid Player ID (-1 for all players)
 * @return Object ID or 0 on failure
 */
int streamer_createDynamicObject(int modelid, const float x, const float y, const float z, const float rx, const float ry, const float rz, int worldid, int interiorid, int playerid);

/**
 * Destroys a dynamic object
 * 
 * @param objectid ID of the object to destroy
 */
void streamer_destroyDynamicObject(int objectid);

#ifdef __cplusplus
}
#endif