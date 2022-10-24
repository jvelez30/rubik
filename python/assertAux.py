from collections import OrderedDict
#Asserts Auxiliar Structures and values
red = "\x1b[31m";
green = "\x1b[32m";
yellow = "\x1b[33m";
blue = "\x1b[34m";
magenta = "\x1b[95m";
white = "\x1b[97m";
positions = OrderedDict([(0,'B1'),(1,'B2'),(2,'B3'),(3,'B4'),(4,'B5'),(5,'B6'),(6,'B7'),(7,'B8'),(8,'B9'),(9,'L1'),(10,'L2'),(11,'L3'),(12,'L4'),(13,'L5'),(14,'L6'),(15,'L7'),(16,'L8'),(17,'L9'),(18,'U1'),(19,'U2'),(20,'U3'),(21,'U4'),(22,'U5'),(23,'U6'),(24,'U7'),(25,'U8'),(26,'U9'),(27,'R1'),(28,'R2'),(29,'R3'),(30,'R4'),(31,'R5'),(32,'R6'),(33,'R7'),(34,'R8'),(35,'R9'),(36,'D1'),(37,'D2'),(38,'D3'),(39,'D4'),(40,'D5'),(41,'D6'),(42,'D7'),(43,'D8'),(44,'D9'),(45,'F1'),(46,'F2'),(47,'F3'),(48,'F4'),(49,'F5'),(50,'F6'),(51,'F7'),(52,'F8'),(53,'F9')]);
orderedValues = OrderedDict([(0,'B1'),(1,'B2'),(2,'B3'),(3,'B4'),(4,'B5'),(5,'B6'),(6,'B7'),(7,'B8'),(8,'B9'),(9,'L1'),(10,'L2'),(11,'L3'),(12,'L4'),(13,'L5'),(14,'L6'),(15,'L7'),(16,'L8'),(17,'L9'),(18,'U1'),(19,'U2'),(20,'U3'),(21,'U4'),(22,'U5'),(23,'U6'),(24,'U7'),(25,'U8'),(26,'U9'),(27,'R1'),(28,'R2'),(29,'R3'),(30,'R4'),(31,'R5'),(32,'R6'),(33,'R7'),(34,'R8'),(35,'R9'),(36,'D1'),(37,'D2'),(38,'D3'),(39,'D4'),(40,'D5'),(41,'D6'),(42,'D7'),(43,'D8'),(44,'D9'),(45,'F1'),(46,'F2'),(47,'F3'),(48,'F4'),(49,'F5'),(50,'F6'),(51,'F7'),(52,'F8'),(53,'F9')]);
orderedColors = OrderedDict([(0,magenta),(1,magenta),(2,magenta),(3,magenta),(4,magenta),(5,magenta),(6,magenta),(7,magenta),(8,magenta),(9,blue),(10,blue),(11,blue),(12,blue),(13,blue),(14,blue),(15,blue),(16,blue),(17,blue),(18,yellow),(19,yellow),(20,yellow),(21,yellow),(22,yellow),(23,yellow),(24,yellow),(25,yellow),(26,yellow),(27,green),(28,green),(29,green),(30,green),(31,green),(32,green),(33,green),(34,green),(35,green),(36,white),(37,white),(38,white),(39,white),(40,white),(41,white),(42,white),(43,white),(44,white),(45,red),(46,red),(47,red),(48,red),(49,red),(50,red),(51,red),(52,red),(53,red)]);
moveUValues = OrderedDict([(0,'B1'),(1,'B2'),(2,'B3'),(3,'B4'),(4,'B5'),(5,'B6'),(6,'L9'),(7,'L6'),(8,'L3'),(9,'L1'),(10,'L2'),(11,'F1'),(12,'L4'),(13,'L5'),(14,'F2'),(15,'L7'),(16,'L8'),(17,'F3'),(18,'U7'),(19,'U4'),(20,'U1'),(21,'U8'),(22,'U5'),(23,'U2'),(24,'U9'),(25,'U6'),(26,'U3'),(27,'B7'),(28,'R2'),(29,'R3'),(30,'B8'),(31,'R5'),(32,'R6'),(33,'B9'),(34,'R8'),(35,'R9'),(36,'D1'),(37,'D2'),(38,'D3'),(39,'D4'),(40,'D5'),(41,'D6'),(42,'D7'),(43,'D8'),(44,'D9'),(45,'R7'),(46,'R4'),(47,'R1'),(48,'F4'),(49,'F5'),(50,'F6'),(51,'F7'),(52,'F8'),(53,'F9')]);
moveUColors = OrderedDict([(0,magenta),(1,magenta),(2,magenta),(3,magenta),(4,magenta),(5,magenta),(6,blue),(7,blue),(8,blue),(9,blue),(10,blue),(11,red),(12,blue),(13,blue),(14,red),(15,blue),(16,blue),(17,red),(18,yellow),(19,yellow),(20,yellow),(21,yellow),(22,yellow),(23,yellow),(24,yellow),(25,yellow),(26,yellow),(27,magenta),(28,green),(29,green),(30,magenta),(31,green),(32,green),(33,magenta),(34,green),(35,green),(36,white),(37,white),(38,white),(39,white),(40,white),(41,white),(42,white),(43,white),(44,white),(45,green),(46,green),(47,green),(48,red),(49,red),(50,red),(51,red),(52,red),(53,red)]);
moveUPValues = OrderedDict([(0,'B1'),(1,'B2'),(2,'B3'),(3,'B4'),(4,'B5'),(5,'B6'),(6,'R1'),(7,'R4'),(8,'R7'),(9,'L1'),(10,'L2'),(11,'B9'),(12,'L4'),(13,'L5'),(14,'B8'),(15,'L7'),(16,'L8'),(17,'B7'),(18,'U3'),(19,'U6'),(20,'U9'),(21,'U2'),(22,'U5'),(23,'U8'),(24,'U1'),(25,'U4'),(26,'U7'),(27,'F3'),(28,'R2'),(29,'R3'),(30,'F2'),(31,'R5'),(32,'R6'),(33,'F1'),(34,'R8'),(35,'R9'),(36,'D1'),(37,'D2'),(38,'D3'),(39,'D4'),(40,'D5'),(41,'D6'),(42,'D7'),(43,'D8'),(44,'D9'),(45,'L3'),(46,'L6'),(47,'L9'),(48,'F4'),(49,'F5'),(50,'F6'),(51,'F7'),(52,'F8'),(53,'F9')]);
moveUPColors = OrderedDict([(0,magenta),(1,magenta),(2,magenta),(3,magenta),(4,magenta),(5,magenta),(6,green),(7,green),(8,green),(9,blue),(10,blue),(11,magenta),(12,blue),(13,blue),(14,magenta),(15,blue),(16,blue),(17,magenta),(18,yellow),(19,yellow),(20,yellow),(21,yellow),(22,yellow),(23,yellow),(24,yellow),(25,yellow),(26,yellow),(27,red),(28,green),(29,green),(30,red),(31,green),(32,green),(33,red),(34,green),(35,green),(36,white),(37,white),(38,white),(39,white),(40,white),(41,white),(42,white),(43,white),(44,white),(45,blue),(46,blue),(47,blue),(48,red),(49,red),(50,red),(51,red),(52,red),(53,red)]);
moveRValues = OrderedDict([(0,'B1'),(1,'B2'),(2,'U3'),(3,'B4'),(4,'B5'),(5,'U6'),(6,'B7'),(7,'B8'),(8,'U9'),(9,'L1'),(10,'L2'),(11,'L3'),(12,'L4'),(13,'L5'),(14,'L6'),(15,'L7'),(16,'L8'),(17,'L9'),(18,'U1'),(19,'U2'),(20,'F3'),(21,'U4'),(22,'U5'),(23,'F6'),(24,'U7'),(25,'U8'),(26,'F9'),(27,'R7'),(28,'R4'),(29,'R1'),(30,'R8'),(31,'R5'),(32,'R2'),(33,'R9'),(34,'R6'),(35,'R3'),(36,'B9'),(37,'D2'),(38,'D3'),(39,'B6'),(40,'D5'),(41,'D6'),(42,'B3'),(43,'D8'),(44,'D9'),(45,'F1'),(46,'F2'),(47,'D7'),(48,'F4'),(49,'F5'),(50,'D4'),(51,'F7'),(52,'F8'),(53,'D1')]);
moveRColors = OrderedDict([(0,magenta),(1,magenta),(2,yellow),(3,magenta),(4,magenta),(5,yellow),(6,magenta),(7,magenta),(8,yellow),(9,blue),(10,blue),(11,blue),(12,blue),(13,blue),(14,blue),(15,blue),(16,blue),(17,blue),(18,yellow),(19,yellow),(20,red),(21,yellow),(22,yellow),(23,red),(24,yellow),(25,yellow),(26,red),(27,green),(28,green),(29,green),(30,green),(31,green),(32,green),(33,green),(34,green),(35,green),(36,magenta),(37,white),(38,white),(39,magenta),(40,white),(41,white),(42,magenta),(43,white),(44,white),(45,red),(46,red),(47,white),(48,red),(49,red),(50,white),(51,red),(52,red),(53,white)]);
moveRPValues = OrderedDict([(0,'B1'),(1,'B2'),(2,'D7'),(3,'B4'),(4,'B5'),(5,'D4'),(6,'B7'),(7,'B8'),(8,'D1'),(9,'L1'),(10,'L2'),(11,'L3'),(12,'L4'),(13,'L5'),(14,'L6'),(15,'L7'),(16,'L8'),(17,'L9'),(18,'U1'),(19,'U2'),(20,'B3'),(21,'U4'),(22,'U5'),(23,'B6'),(24,'U7'),(25,'U8'),(26,'B9'),(27,'R3'),(28,'R6'),(29,'R9'),(30,'R2'),(31,'R5'),(32,'R8'),(33,'R1'),(34,'R4'),(35,'R7'),(36,'F9'),(37,'D2'),(38,'D3'),(39,'F6'),(40,'D5'),(41,'D6'),(42,'F3'),(43,'D8'),(44,'D9'),(45,'F1'),(46,'F2'),(47,'U3'),(48,'F4'),(49,'F5'),(50,'U6'),(51,'F7'),(52,'F8'),(53,'U9')]);
moveRPColors = OrderedDict([(0,magenta),(1,magenta),(2,white),(3,magenta),(4,magenta),(5,white),(6,magenta),(7,magenta),(8,white),(9,blue),(10,blue),(11,blue),(12,blue),(13,blue),(14,blue),(15,blue),(16,blue),(17,blue),(18,yellow),(19,yellow),(20,magenta),(21,yellow),(22,yellow),(23,magenta),(24,yellow),(25,yellow),(26,magenta),(27,green),(28,green),(29,green),(30,green),(31,green),(32,green),(33,green),(34,green),(35,green),(36,red),(37,white),(38,white),(39,red),(40,white),(41,white),(42,red),(43,white),(44,white),(45,red),(46,red),(47,yellow),(48,red),(49,red),(50,yellow),(51,red),(52,red),(53,yellow)]);
moveLValues = OrderedDict([(0,'D9'),(1,'B2'),(2,'B3'),(3,'D6'),(4,'B5'),(5,'B6'),(6,'D3'),(7,'B8'),(8,'B9'),(9,'L7'),(10,'L4'),(11,'L1'),(12,'L8'),(13,'L5'),(14,'L2'),(15,'L9'),(16,'L6'),(17,'L3'),(18,'B1'),(19,'U2'),(20,'U3'),(21,'B4'),(22,'U5'),(23,'U6'),(24,'B7'),(25,'U8'),(26,'U9'),(27,'R1'),(28,'R2'),(29,'R3'),(30,'R4'),(31,'R5'),(32,'R6'),(33,'R7'),(34,'R8'),(35,'R9'),(36,'D1'),(37,'D2'),(38,'F7'),(39,'D4'),(40,'D5'),(41,'F4'),(42,'D7'),(43,'D8'),(44,'F1'),(45,'U1'),(46,'F2'),(47,'F3'),(48,'U4'),(49,'F5'),(50,'F6'),(51,'U7'),(52,'F8'),(53,'F9')]);
moveLColors = OrderedDict([(0,white),(1,magenta),(2,magenta),(3,white),(4,magenta),(5,magenta),(6,white),(7,magenta),(8,magenta),(9,blue),(10,blue),(11,blue),(12,blue),(13,blue),(14,blue),(15,blue),(16,blue),(17,blue),(18,magenta),(19,yellow),(20,yellow),(21,magenta),(22,yellow),(23,yellow),(24,magenta),(25,yellow),(26,yellow),(27,green),(28,green),(29,green),(30,green),(31,green),(32,green),(33,green),(34,green),(35,green),(36,white),(37,white),(38,red),(39,white),(40,white),(41,red),(42,white),(43,white),(44,red),(45,yellow),(46,red),(47,red),(48,yellow),(49,red),(50,red),(51,yellow),(52,red),(53,red)]);
moveLPValues = OrderedDict([(0,'U1'),(1,'B2'),(2,'B3'),(3,'U4'),(4,'B5'),(5,'B6'),(6,'U7'),(7,'B8'),(8,'B9'),(9,'L3'),(10,'L6'),(11,'L9'),(12,'L2'),(13,'L5'),(14,'L8'),(15,'L1'),(16,'L4'),(17,'L7'),(18,'F1'),(19,'U2'),(20,'U3'),(21,'F4'),(22,'U5'),(23,'U6'),(24,'F7'),(25,'U8'),(26,'U9'),(27,'R1'),(28,'R2'),(29,'R3'),(30,'R4'),(31,'R5'),(32,'R6'),(33,'R7'),(34,'R8'),(35,'R9'),(36,'D1'),(37,'D2'),(38,'B7'),(39,'D4'),(40,'D5'),(41,'B4'),(42,'D7'),(43,'D8'),(44,'B1'),(45,'D9'),(46,'F2'),(47,'F3'),(48,'D6'),(49,'F5'),(50,'F6'),(51,'D3'),(52,'F8'),(53,'F9')]);
moveLPColors = OrderedDict([(0,yellow),(1,magenta),(2,magenta),(3,yellow),(4,magenta),(5,magenta),(6,yellow),(7,magenta),(8,magenta),(9,blue),(10,blue),(11,blue),(12,blue),(13,blue),(14,blue),(15,blue),(16,blue),(17,blue),(18,red),(19,yellow),(20,yellow),(21,red),(22,yellow),(23,yellow),(24,red),(25,yellow),(26,yellow),(27,green),(28,green),(29,green),(30,green),(31,green),(32,green),(33,green),(34,green),(35,green),(36,white),(37,white),(38,magenta),(39,white),(40,white),(41,magenta),(42,white),(43,white),(44,magenta),(45,white),(46,red),(47,red),(48,white),(49,red),(50,red),(51,white),(52,red),(53,red)]);
moveFValues = OrderedDict([(0,'B1'),(1,'B2'),(2,'B3'),(3,'B4'),(4,'B5'),(5,'B6'),(6,'B7'),(7,'B8'),(8,'B9'),(9,'L1'),(10,'L2'),(11,'L3'),(12,'L4'),(13,'L5'),(14,'L6'),(15,'D7'),(16,'D8'),(17,'D9'),(18,'U1'),(19,'U2'),(20,'U3'),(21,'U4'),(22,'U5'),(23,'U6'),(24,'L7'),(25,'L8'),(26,'L9'),(27,'R1'),(28,'R2'),(29,'R3'),(30,'R4'),(31,'R5'),(32,'R6'),(33,'U7'),(34,'U8'),(35,'U9'),(36,'D1'),(37,'D2'),(38,'D3'),(39,'D4'),(40,'D5'),(41,'D6'),(42,'R7'),(43,'R8'),(44,'R9'),(45,'F7'),(46,'F4'),(47,'F1'),(48,'F8'),(49,'F5'),(50,'F2'),(51,'F9'),(52,'F6'),(53,'F3')]);
moveFColors = OrderedDict([(0,magenta),(1,magenta),(2,magenta),(3,magenta),(4,magenta),(5,magenta),(6,magenta),(7,magenta),(8,magenta),(9,blue),(10,blue),(11,blue),(12,blue),(13,blue),(14,blue),(15,white),(16,white),(17,white),(18,yellow),(19,yellow),(20,yellow),(21,yellow),(22,yellow),(23,yellow),(24,blue),(25,blue),(26,blue),(27,green),(28,green),(29,green),(30,green),(31,green),(32,green),(33,yellow),(34,yellow),(35,yellow),(36,white),(37,white),(38,white),(39,white),(40,white),(41,white),(42,green),(43,green),(44,green),(45,red),(46,red),(47,red),(48,red),(49,red),(50,red),(51,red),(52,red),(53,red)]);
moveFPValues = OrderedDict([(0,'B1'),(1,'B2'),(2,'B3'),(3,'B4'),(4,'B5'),(5,'B6'),(6,'B7'),(7,'B8'),(8,'B9'),(9,'L1'),(10,'L2'),(11,'L3'),(12,'L4'),(13,'L5'),(14,'L6'),(15,'U7'),(16,'U8'),(17,'U9'),(18,'U1'),(19,'U2'),(20,'U3'),(21,'U4'),(22,'U5'),(23,'U6'),(24,'R7'),(25,'R8'),(26,'R9'),(27,'R1'),(28,'R2'),(29,'R3'),(30,'R4'),(31,'R5'),(32,'R6'),(33,'D7'),(34,'D8'),(35,'D9'),(36,'D1'),(37,'D2'),(38,'D3'),(39,'D4'),(40,'D5'),(41,'D6'),(42,'L7'),(43,'L8'),(44,'L9'),(45,'F3'),(46,'F6'),(47,'F9'),(48,'F2'),(49,'F5'),(50,'F8'),(51,'F1'),(52,'F4'),(53,'F7')]);
moveFPColors = OrderedDict([(0,magenta),(1,magenta),(2,magenta),(3,magenta),(4,magenta),(5,magenta),(6,magenta),(7,magenta),(8,magenta),(9,blue),(10,blue),(11,blue),(12,blue),(13,blue),(14,blue),(15,yellow),(16,yellow),(17,yellow),(18,yellow),(19,yellow),(20,yellow),(21,yellow),(22,yellow),(23,yellow),(24,green),(25,green),(26,green),(27,green),(28,green),(29,green),(30,green),(31,green),(32,green),(33,white),(34,white),(35,white),(36,white),(37,white),(38,white),(39,white),(40,white),(41,white),(42,blue),(43,blue),(44,blue),(45,red),(46,red),(47,red),(48,red),(49,red),(50,red),(51,red),(52,red),(53,red)]);
moveBValues = OrderedDict([(0,'B7'),(1,'B4'),(2,'B1'),(3,'B8'),(4,'B5'),(5,'B2'),(6,'B9'),(7,'B6'),(8,'B3'),(9,'U1'),(10,'U2'),(11,'U3'),(12,'L4'),(13,'L5'),(14,'L6'),(15,'L7'),(16,'L8'),(17,'L9'),(18,'R1'),(19,'R2'),(20,'R3'),(21,'U4'),(22,'U5'),(23,'U6'),(24,'U7'),(25,'U8'),(26,'U9'),(27,'D1'),(28,'D2'),(29,'D3'),(30,'R4'),(31,'R5'),(32,'R6'),(33,'R7'),(34,'R8'),(35,'R9'),(36,'L1'),(37,'L2'),(38,'L3'),(39,'D4'),(40,'D5'),(41,'D6'),(42,'D7'),(43,'D8'),(44,'D9'),(45,'F1'),(46,'F2'),(47,'F3'),(48,'F4'),(49,'F5'),(50,'F6'),(51,'F7'),(52,'F8'),(53,'F9')]);
moveBColors = OrderedDict([(0,magenta),(1,magenta),(2,magenta),(3,magenta),(4,magenta),(5,magenta),(6,magenta),(7,magenta),(8,magenta),(9,yellow),(10,yellow),(11,yellow),(12,blue),(13,blue),(14,blue),(15,blue),(16,blue),(17,blue),(18,green),(19,green),(20,green),(21,yellow),(22,yellow),(23,yellow),(24,yellow),(25,yellow),(26,yellow),(27,white),(28,white),(29,white),(30,green),(31,green),(32,green),(33,green),(34,green),(35,green),(36,blue),(37,blue),(38,blue),(39,white),(40,white),(41,white),(42,white),(43,white),(44,white),(45,red),(46,red),(47,red),(48,red),(49,red),(50,red),(51,red),(52,red),(53,red)]);
moveBPValues = OrderedDict([(0,'B3'),(1,'B6'),(2,'B9'),(3,'B2'),(4,'B5'),(5,'B8'),(6,'B1'),(7,'B4'),(8,'B7'),(9,'D1'),(10,'D2'),(11,'D3'),(12,'L4'),(13,'L5'),(14,'L6'),(15,'L7'),(16,'L8'),(17,'L9'),(18,'L1'),(19,'L2'),(20,'L3'),(21,'U4'),(22,'U5'),(23,'U6'),(24,'U7'),(25,'U8'),(26,'U9'),(27,'U1'),(28,'U2'),(29,'U3'),(30,'R4'),(31,'R5'),(32,'R6'),(33,'R7'),(34,'R8'),(35,'R9'),(36,'R1'),(37,'R2'),(38,'R3'),(39,'D4'),(40,'D5'),(41,'D6'),(42,'D7'),(43,'D8'),(44,'D9'),(45,'F1'),(46,'F2'),(47,'F3'),(48,'F4'),(49,'F5'),(50,'F6'),(51,'F7'),(52,'F8'),(53,'F9')]);
moveBPColors = OrderedDict([(0,magenta),(1,magenta),(2,magenta),(3,magenta),(4,magenta),(5,magenta),(6,magenta),(7,magenta),(8,magenta),(9,white),(10,white),(11,white),(12,blue),(13,blue),(14,blue),(15,blue),(16,blue),(17,blue),(18,blue),(19,blue),(20,blue),(21,yellow),(22,yellow),(23,yellow),(24,yellow),(25,yellow),(26,yellow),(27,yellow),(28,yellow),(29,yellow),(30,green),(31,green),(32,green),(33,green),(34,green),(35,green),(36,green),(37,green),(38,green),(39,white),(40,white),(41,white),(42,white),(43,white),(44,white),(45,red),(46,red),(47,red),(48,red),(49,red),(50,red),(51,red),(52,red),(53,red)]);
moveDValues = OrderedDict([(0,'R3'),(1,'R6'),(2,'R9'),(3,'B4'),(4,'B5'),(5,'B6'),(6,'B7'),(7,'B8'),(8,'B9'),(9,'B3'),(10,'L2'),(11,'L3'),(12,'B2'),(13,'L5'),(14,'L6'),(15,'B1'),(16,'L8'),(17,'L9'),(18,'U1'),(19,'U2'),(20,'U3'),(21,'U4'),(22,'U5'),(23,'U6'),(24,'U7'),(25,'U8'),(26,'U9'),(27,'R1'),(28,'R2'),(29,'F9'),(30,'R4'),(31,'R5'),(32,'F8'),(33,'R7'),(34,'R8'),(35,'F7'),(36,'D7'),(37,'D4'),(38,'D1'),(39,'D8'),(40,'D5'),(41,'D2'),(42,'D9'),(43,'D6'),(44,'D3'),(45,'F1'),(46,'F2'),(47,'F3'),(48,'F4'),(49,'F5'),(50,'F6'),(51,'L1'),(52,'L4'),(53,'L7')]);
moveDColors = OrderedDict([(0,green),(1,green),(2,green),(3,magenta),(4,magenta),(5,magenta),(6,magenta),(7,magenta),(8,magenta),(9,magenta),(10,blue),(11,blue),(12,magenta),(13,blue),(14,blue),(15,magenta),(16,blue),(17,blue),(18,yellow),(19,yellow),(20,yellow),(21,yellow),(22,yellow),(23,yellow),(24,yellow),(25,yellow),(26,yellow),(27,green),(28,green),(29,red),(30,green),(31,green),(32,red),(33,green),(34,green),(35,red),(36,white),(37,white),(38,white),(39,white),(40,white),(41,white),(42,white),(43,white),(44,white),(45,red),(46,red),(47,red),(48,red),(49,red),(50,red),(51,blue),(52,blue),(53,blue)]);
moveDPValues = OrderedDict([(0,'L7'),(1,'L4'),(2,'L1'),(3,'B4'),(4,'B5'),(5,'B6'),(6,'B7'),(7,'B8'),(8,'B9'),(9,'F7'),(10,'L2'),(11,'L3'),(12,'F8'),(13,'L5'),(14,'L6'),(15,'F9'),(16,'L8'),(17,'L9'),(18,'U1'),(19,'U2'),(20,'U3'),(21,'U4'),(22,'U5'),(23,'U6'),(24,'U7'),(25,'U8'),(26,'U9'),(27,'R1'),(28,'R2'),(29,'B1'),(30,'R4'),(31,'R5'),(32,'B2'),(33,'R7'),(34,'R8'),(35,'B3'),(36,'D3'),(37,'D6'),(38,'D9'),(39,'D2'),(40,'D5'),(41,'D8'),(42,'D1'),(43,'D4'),(44,'D7'),(45,'F1'),(46,'F2'),(47,'F3'),(48,'F4'),(49,'F5'),(50,'F6'),(51,'R9'),(52,'R6'),(53,'R3')]);
moveDPColors = OrderedDict([(0,blue),(1,blue),(2,blue),(3,magenta),(4,magenta),(5,magenta),(6,magenta),(7,magenta),(8,magenta),(9,red),(10,blue),(11,blue),(12,red),(13,blue),(14,blue),(15,red),(16,blue),(17,blue),(18,yellow),(19,yellow),(20,yellow),(21,yellow),(22,yellow),(23,yellow),(24,yellow),(25,yellow),(26,yellow),(27,green),(28,green),(29,magenta),(30,green),(31,green),(32,magenta),(33,green),(34,green),(35,magenta),(36,white),(37,white),(38,white),(39,white),(40,white),(41,white),(42,white),(43,white),(44,white),(45,red),(46,red),(47,red),(48,red),(49,red),(50,red),(51,green),(52,green),(53,green)]);
sexyMoveValues = OrderedDict([(0,'B1'),(1,'B2'),(2,'B3'),(3,'B4'),(4,'B5'),(5,'B6'),(6,'R1'),(7,'R4'),(8,'B7'),(9,'L1'),(10,'L2'),(11,'B9'),(12,'L4'),(13,'L5'),(14,'L6'),(15,'L7'),(16,'L8'),(17,'L9'),(18,'U3'),(19,'U6'),(20,'L3'),(21,'U4'),(22,'U5'),(23,'F6'),(24,'U7'),(25,'U8'),(26,'F9'),(27,'U1'),(28,'R2'),(29,'R3'),(30,'R8'),(31,'R5'),(32,'R6'),(33,'R9'),(34,'B8'),(35,'U9'),(36,'D1'),(37,'D2'),(38,'D3'),(39,'D4'),(40,'D5'),(41,'D6'),(42,'R7'),(43,'D8'),(44,'D9'),(45,'F1'),(46,'F2'),(47,'D7'),(48,'F4'),(49,'F5'),(50,'U2'),(51,'F7'),(52,'F8'),(53,'F3')]);
sexyMoveColors = OrderedDict([(0,magenta),(1,magenta),(2,magenta),(3,magenta),(4,magenta),(5,magenta),(6,green),(7,green),(8,magenta),(9,blue),(10,blue),(11,magenta),(12,blue),(13,blue),(14,blue),(15,blue),(16,blue),(17,blue),(18,yellow),(19,yellow),(20,blue),(21,yellow),(22,yellow),(23,red),(24,yellow),(25,yellow),(26,red),(27,yellow),(28,green),(29,green),(30,green),(31,green),(32,green),(33,green),(34,magenta),(35,yellow),(36,white),(37,white),(38,white),(39,white),(40,white),(41,white),(42,green),(43,white),(44,white),(45,red),(46,red),(47,white),(48,red),(49,red),(50,yellow),(51,red),(52,red),(53,red)]);